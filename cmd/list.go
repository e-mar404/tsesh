package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/spf13/cobra"
)

var nameStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("62")).
	Foreground(lipgloss.Color("230"))

var urlStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

var contentStyle = lipgloss.NewStyle().
	PaddingLeft(1)

var itemStyle = lipgloss.NewStyle().
	PaddingBottom(1)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list all bookmarks set for current directory",
		PreRunE: batch(
			bookmark.ValidateDataStorage,
			data.Load,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := data.List()
			if err != nil {
				return err
			}

			if len(list) == 0 {
				log.Print("no bookmarks set for this directory")
				return nil
			}

			out := strings.Builder{}
			for i, mark := range list {
				content := lipgloss.JoinVertical(lipgloss.Left,
					nameStyle.Render(mark.Name),
					urlStyle.Render(mark.Url),
				)

				item := lipgloss.JoinHorizontal(lipgloss.Top,
					fmt.Sprintf("[%d]", i),
					contentStyle.Render(content),
				)

				str := fmt.Sprintf("%s\n", itemStyle.Render(item))
				out.WriteString(str)
			}

			fmt.Print(out.String())

			return nil
		},
	}
)
