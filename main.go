package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	path string
	name string
}

func (i item) FilterValue() string { return i.path }

func (i item) Title() string { return i.path }

func (i item) Description() string { return "" }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			selected := m.list.SelectedItem().(item)
			handleTmuxSession(selected.path)
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func findDirectories() ([]item, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dirs := []string{home, filepath.Join(home, "projects")}
	var items []item

	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			name := entry.Name()
			if strings.HasPrefix(name, ".") && name != ".dotfiles" {
				continue
			}

			path := filepath.Join(dir, name)
			items = append(items, item{path: path, name: name})
		}
	}

	return items, nil
}

func sanitizeSessionName(name string) string {
	replacer := strings.NewReplacer(" ", "_", "-", "_", ".", "_")
	return replacer.Replace(name)
}

func isTmuxRunning() bool {
	cmd := exec.Command("pgrep", "tmux")
	err := cmd.Run()
	return err == nil
}

func hasSession(name string) bool {
	cmd := exec.Command("tmux", "has-session", "-t="+name)
	err := cmd.Run()
	return err == nil
}

func handleTmuxSession(selected string) {
	selectedName := sanitizeSessionName(filepath.Base(selected))
	tmuxRunning := isTmuxRunning()
	inTmux := os.Getenv("TMUX") != ""

	if !inTmux && !tmuxRunning {
		cmd := exec.Command("tmux", "new-session", "-s", selectedName, "-c", selected)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return
	}

	if !hasSession(selectedName) {
		exec.Command("tmux", "new-session", "-ds", selectedName, "-c", selected).Run()
	}

	if inTmux && tmuxRunning {
		exec.Command("tmux", "switch-client", "-t", selectedName).Run()
		return
	}

	cmd := exec.Command("tmux", "attach-session", "-t", selectedName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	if len(os.Args) == 2 {
		handleTmuxSession(os.Args[1])
		return
	}

	items, err := findDirectories()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding directories: %v\n", err)
		os.Exit(1)
	}

	if len(items) == 0 {
		fmt.Println("No directories found")
		os.Exit(0)
	}

	delegate := list.NewDefaultDelegate()
	listItems := make([]list.Item, len(items))
	for i, it := range items {
		listItems[i] = it
	}
	l := list.New(listItems, delegate, 0, 0)
	l.Title = "Select a project"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1)

	m := model{list: l}
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
