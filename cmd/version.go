package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "get tsesh cli version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("tsesh version %s\n", version)
		},
	}
)
