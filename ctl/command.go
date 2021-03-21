package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "leetcode-go",
	Short: "A simple command line client for leetcode-go.",
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	rootCmd.AddCommand(
		versionCmd,
		newBuildCommand(),
		newLabelCommand(),
		newPDFCommand(),
		newRefresh(),
	)
}
