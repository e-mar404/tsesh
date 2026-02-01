package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/e-mar404/tsesh/picker"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "tsesh",
	Short: "terminal sessionizer extending tmux",
	Run: func (cmd *cobra.Command, args []string) {
		p := tea.NewProgram(picker.New(), tea.WithAltScreen())
		if pi, err := p.Run(); err != nil {
			fmt.Printf("%v\n", pi.(picker.Picker).Err)
			fmt.Printf("Encountered an error when trying to run the directory picker: %v\n", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("could not execute command: %v\n", err)
		os.Exit(1)
	}
}

