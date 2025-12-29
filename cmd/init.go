package cmd

import (
	"github.com/andre/forge/internal/scaffold"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project]",
	Short: "Initialize a new Go backend project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return scaffold.InitProject(args[0])
	},
}
