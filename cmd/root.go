package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/e-mar404/tsesh/internal/config"
	"github.com/e-mar404/tsesh/internal/picker"
	"github.com/spf13/cobra"
)

var (
	version string
	cfg     = &config.Config{}
	data    = make(bookmark.Data)
	rootCmd = &cobra.Command{
		Use:   "tsesh",
		Short: "terminal sessionizer extending tmux",
		RunE: func(cmd *cobra.Command, args []string) error {
			p := tea.NewProgram(picker.New(cfg), tea.WithAltScreen())
			if pi, err := p.Run(); err != nil {
				fmt.Printf("%v\n", pi.(picker.Picker).Err)
				return fmt.Errorf("Encountered an error when trying to run the directory picker: %v\n", err)
			}
			return nil
		},
	}
)

func Execute(currentVersion string) {
	version = currentVersion
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(loadConfig)

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(versionCmd)
}

func loadConfig() {
	if !config.Exists() {
		err := config.CreateDefault()
		cobra.CheckErr(err)
	}

	err := config.LoadInto(cfg)
	cobra.CheckErr(err)
}

func batch(checks ...func() error) func(*cobra.Command, []string) error {
	return func(_ *cobra.Command, _ []string) error {
		for _, check := range checks {
			if err := check(); err != nil {
				return err
			}
		}
		return nil
	}
}
