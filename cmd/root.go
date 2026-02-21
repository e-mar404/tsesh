package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/e-mar404/tsesh/config"
	"github.com/e-mar404/tsesh/picker"
	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
	rootCmd = &cobra.Command{
		Use:   "tsesh",
		Short: "terminal sessionizer extending tmux",
		Run: func(cmd *cobra.Command, args []string) {
			searchPaths := []string{
				"~/",
				"~/projects",
			}

			p := tea.NewProgram(picker.New(searchPaths), tea.WithAltScreen())
			if pi, err := p.Run(); err != nil {
				fmt.Printf("%v\n", pi.(picker.Picker).Err)
				fmt.Printf("Encountered an error when trying to run the directory picker: %v\n", err)
				os.Exit(1)
			}
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	_, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// check for config file 
	// if it doesnt exist then make a default with config.CreateDefault()
	// else load config file and save into cfg *Config
}
