package cmd

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/spf13/cobra"
)

var (
	url    string
	paste  bool
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a directory or a url to the current working directory",
		Args:  cobra.NoArgs,
		PreRunE: batch(
			bookmark.ValidateDataStorage,
			data.Load,
		),
		Run: func(cmd *cobra.Command, args []string) {
			if paste {
				seshType := os.Getenv("XDG_SESSION_TYPE")
				if seshType != "wayland" {
					log.Fatal("pasting from clipboard is not currently available for non wayland devices")
				}

				output, err := exec.Command("wl-paste").Output()
				if err != nil {
					cobra.CheckErr(err)
				}

				url = string(output)
			}

			err := data.Add(name, url)
			cobra.CheckErr(err)

			err = data.Save()
			cobra.CheckErr(err)
		},
	}
)

func init() {
	addCmd.Flags().StringVarP(&name, "name", "n", "", "name for bookmark")
	addCmd.Flags().StringVarP(&url, "url", "u", "", "url for bookmark")
	addCmd.Flags().BoolVarP(&paste, "paste", "p", false, "use current clipboard as url")

	addCmd.MarkFlagsOneRequired("url", "paste")
}
