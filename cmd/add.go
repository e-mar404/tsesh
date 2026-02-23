package cmd

import (
	"fmt"
	"os"

	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/spf13/cobra"
)

var (
	data   = make(bookmark.Data)
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a directory or a url to the current working directory",
		PreRunE: batch(
			bookmark.ValidateDataStorage,
			data.Load,
		),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("command needs at least 1 argument. Add either a url or directory path")
				os.Exit(1)
			}

			err := data.Add(args...)
			cobra.CheckErr(err)
		},
	}
)
