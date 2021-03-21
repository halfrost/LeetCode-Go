package main

import (
	"github.com/spf13/cobra"
)

func newRefresh() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Refresh all document",
		Run: func(cmd *cobra.Command, args []string) {
			refresh()
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func refresh() {
	//buildBookMenu()
	copyLackFile()
	delPreNext()
	buildREADME()
	buildChapterTwo(true)
	addPreNext()
}
