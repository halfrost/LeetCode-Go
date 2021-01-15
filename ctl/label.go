package main

import (
	"github.com/spf13/cobra"
)

func newLabelCommand() *cobra.Command {
	mc := &cobra.Command{
		Use:   "label <subcommand>",
		Short: "Label related commands",
	}
	//mc.PersistentFlags().StringVar(&logicEndpoint, "endpoint", "localhost:5880", "endpoint of logic service")
	mc.AddCommand(
		newAddPreNext(),
		newDeletePreNext(),
	)
	return mc
}

func newAddPreNext() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-pre-next",
		Short: "Add pre-next label",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func newDeletePreNext() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-pre-next",
		Short: "Delete pre-next label",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}
