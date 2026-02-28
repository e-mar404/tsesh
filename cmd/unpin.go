package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	unpinCmd = &cobra.Command{
		Use:   "unpin",
		Short: "unpin directory from fuzzy finder if it is in the pinned list",
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			return cfg.Unpin(cwd)
		},
	}
)
