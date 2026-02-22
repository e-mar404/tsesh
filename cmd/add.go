package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"github.com/e-mar404/tsesh/internal/bookmark"
	"github.com/spf13/cobra"
)

var (
	data   = bookmark.Entries{}
	addCmd = &cobra.Command{
		Use:     "add",
		Short:   "add a directory or a url to the current working directory",
		PreRunE: validateDataFile,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("command needs at least 1 argument. Add either a url or directory path")
				os.Exit(1)
			}

			cwd, _ := os.Getwd()
			for _, arg := range args {
				fmt.Printf("adding url: %v to %v\n", arg, cwd)
			}
		},
	}
)

func validateDataFile(_ *cobra.Command, _ []string) error {
	path, err := xdg.DataFile("tsesh/data.json")
	if err != nil {
		return err
	}

	_, err = os.Stat(path)
	if err == nil {
		return loadDataFile(path)
	}

	buf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return os.WriteFile(path, buf.Bytes(), os.ModePerm)
}

func loadDataFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(f)
	return decoder.Decode(&data)
}
