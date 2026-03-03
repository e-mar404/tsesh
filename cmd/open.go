package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/spf13/cobra"
)

var (
	openAll bool
	openCmd = &cobra.Command{
		Use:   "open",
		Short: "open a bookmark by index",
		PreRunE: batch(
			bookmark.ValidateDataStorage,
			data.Load,
		),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if openAll {
				if id != -1 || name != "" {
					log.Warn("flag -a enabled, opening all bookmarks and ignoring any other flags")
				}

				return data.OpenAll()
			}

			if id != -1 {
				return data.OpenIndex(id)
			}

			return data.OpenName(name)
		},
	}
)

func init() {
	openCmd.Flags().BoolVarP(&openAll, "all", "a", false, "open all bookmarks for this directory")
	openCmd.Flags().StringVarP(&name, "name", "n", "", "open bookmark by name")
	openCmd.Flags().IntVar(&id, "id", -1, "open bookmark by id")
	openCmd.MarkFlagsMutuallyExclusive("name", "id")
	openCmd.MarkFlagsOneRequired("all", "name", "id")
}
