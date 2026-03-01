package cmd

import (
	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a directory or a url to the current working directory",
		Args:  cobra.RangeArgs(1, 10),
		PreRunE: batch(
			bookmark.ValidateDataStorage,
			data.Load,
		),
		Run: func(cmd *cobra.Command, args []string) {
			err := data.Add(args...)
			cobra.CheckErr(err)

			err = data.Save()
			cobra.CheckErr(err)
		},
	}
)
