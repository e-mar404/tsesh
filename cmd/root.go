package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/e-mar404/tsesh/internal/config"
	"github.com/e-mar404/tsesh/internal/picker"
	"github.com/spf13/cobra"
)

var (
	cfg     = &config.Config{}
	rootCmd = &cobra.Command{
		Use:   "tsesh",
		Short: "terminal sessionizer extending tmux",
		Run: func(cmd *cobra.Command, args []string) {
			p := tea.NewProgram(picker.New(cfg), tea.WithAltScreen())
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
	cobra.OnInitialize(loadConfig)

	rootCmd.AddCommand(addCmd)
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
