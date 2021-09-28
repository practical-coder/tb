package cmd

import (
	"tb/which"

	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "which functionality - find absolute paths of executables in arguments",
	Run: func(cmd *cobra.Command, args []string) {
		which.Find(args...)
	},
}
