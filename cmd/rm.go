package cmd

import (
	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/spf13/cobra"
)

var (
	rmCmd = &cobra.Command{
		Use:   "rm",
		Short: "remove url provided from bookmark list for current working directory",
		Args:  cobra.ExactArgs(1),
		PreRunE: batch(
			bookmark.ValidateDataStorage,
			data.Load,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := data.Remove(args[0]); err != nil {
				return err
			}

			return data.Save()
		},
	}
)
