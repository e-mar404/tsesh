package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	pinCmd = &cobra.Command{
		Use:   "pin",
		Short: "pins the current directory to the top of the list when running tsesh",
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			return cfg.Pin(cwd)
		},
	}
)
