package cmd

import (
	"strconv"

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
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if openAll {
				if len(args) > 0 {
					log.Warn("flag -a enabled, opening all bookmarks and ignoring arguments passed")
				}

				return data.OpenAll()
			}

			idx, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			return data.Open(idx)
		},
	}
)

func init() {
	openCmd.Flags().BoolVarP(&openAll, "all", "a", false, "open all bookmarks for this directory")
}
