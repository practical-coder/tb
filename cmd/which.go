package cmd

import (
	"tb/which"

	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "search $PATH for executables",
	Long: "search $PATH for executables, print absolute paths",
	Run: func(cmd *cobra.Command, args []string) {
		which.Find(args...)
	},
}
