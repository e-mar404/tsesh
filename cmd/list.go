package cmd

import (
	"fmt"

	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list all bookmarks set for current directory",
		PreRunE: batch(
			bookmark.ValidateDataStorage,
			data.Load,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := data.List()
			if err != nil {
				return err
			}

			if len(list) == 0 {
				fmt.Printf("no bookmarks set for this directory\n")
				return nil
			}

			for i, mark := range list {
				fmt.Printf("[%v] %v\n", i, mark.Url)
			}

			return nil
		},
	}
)
