package cmd

import (
	"github.com/drewjya/forge/internal/scaffold"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Generate a new application module",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return scaffold.AddModule(args[0])
	},
}
