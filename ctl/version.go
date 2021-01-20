package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version    = "v1.0"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version of tacoctl",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("tacoctl version:", version)
		},
	}
)
