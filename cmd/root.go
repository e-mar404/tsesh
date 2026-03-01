package cmd

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/e-mar404/tsesh/internal/config"
	"github.com/e-mar404/tsesh/internal/picker"
	"github.com/spf13/cobra"
)

var (
	version string
	verbose bool
	cfg     = &config.Config{}
	data    = make(bookmark.Data)
	rootCmd = &cobra.Command{
		Use:              "tsesh",
		Short:            "terminal sessionizer extending tmux",
		PersistentPreRun: setLogLevel,
		RunE: func(cmd *cobra.Command, args []string) error {
			p := tea.NewProgram(picker.New(cfg), tea.WithAltScreen())
			if pi, err := p.Run(); err != nil {
				log.Errorf("%v\n", pi.(picker.Picker).Err)
				os.Exit(1)
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
	rootCmd.AddCommand(pinCmd)
	rootCmd.AddCommand(unpinCmd)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "show debug logging")
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

func setLogLevel(_ *cobra.Command, _ []string) {
	level := log.Level(0)
	if verbose {
		log.Info("verbose flag set")
		level = log.Level(-4)
	}
	log.SetLevel(level)
	log.SetOutput(os.Stderr)
	log.SetReportTimestamp(false)
}
