package cmd

import (
	"github.com/drewjya/forge/internal/scaffold"
	"github.com/spf13/cobra"
)

var removeModuleCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove an application module",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return scaffold.RemoveModule(args[0])
	},
}

func init() {
	rootCmd.AddCommand(removeModuleCmd)
}
