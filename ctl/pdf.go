package main

import (
	"github.com/spf13/cobra"
)

func newPDFCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pdf <subcommand>",
		Short: "PDF related commands",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	//mc.PersistentFlags().StringVar(&logicEndpoint, "endpoint", "localhost:5880", "endpoint of logic service")
	return cmd
}
