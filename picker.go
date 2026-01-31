package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/e-mar404/tsesh/tmux"
)

// Wrapper for list.Item to add extra fields
type item struct {
	name string
	path string
	// TODO: at some point I should also add a configuration var for code that should be executed before and after entering the tmux session
}

func (i item) FilterValue() string {
	return i.name
}

func (i item) Title() string {
	return i.name
}

func (i item) Description() string {
	return i.path
}

type picker struct {
	list list.Model
}

func (p picker) Init() tea.Cmd {
	return nil
}

func (p picker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.list.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return p, tea.Quit

		case "enter":
			if p.list.SelectedItem() == nil {
				return p, tea.Quit
			}

			choice := p.list.SelectedItem().(item)
			switch tmux.HasSession(choice.name) {
			case true:
				if tmux.Inside() {
					tmux.SwitchClient(choice.name)
					return p, tea.Quit
				}
				tmux.Attach(choice.name)

			case false:
				if tmux.Inside() {
					tmux.NewSession(choice.name, choice.path, true)
					tmux.SwitchClient(choice.name)
					return p, tea.Quit
				}
				tmux.NewSession(choice.name, choice.path, false)
			}
			return p, tea.Quit
		}
	}

	p.list, cmd = p.list.Update(msg)

	return p, cmd
}

func (p picker) View() string {
	return p.list.View()
}

func newPicker() picker {
	// TODO: this is were I should load the files from the config
	// for now they will be hardcoded
	searchPaths := []string{
		"~/",
		"~/code",
		"~/projects",
	}

	return picker{
		list: list.New(
			findDirectories(searchPaths),
			list.NewDefaultDelegate(),
			0,
			0,
		),
	}
}
