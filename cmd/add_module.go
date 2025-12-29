package cmd

import (
	"github.com/drewjya/forge/internal/scaffold"
	"github.com/spf13/cobra"
)

var noCrud bool

var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new application module",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return scaffold.AddModule(args[0], !noCrud)
	},
}

func init() {
	addCmd.Flags().BoolVar(&noCrud, "no-crud", false, "Skip CRUD generation")
	rootCmd.AddCommand(addCmd)
}
