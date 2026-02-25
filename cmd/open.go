package cmd

import (
	"strconv"

	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/spf13/cobra"
)

var (
	openCmd = &cobra.Command{
		Use:   "open",
		Short: "open a bookmark by index",
		PreRunE: batch(
			bookmark.ValidateDataStorage,
			data.Load,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			idx, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			return data.Open(idx)
		},
	}
)
