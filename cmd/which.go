package cmd

import (
	"tb/which"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "search $PATH for executables",
	Long:  "search $PATH for executables, print absolute paths",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Info().Msgf("which requires at least one executable as argument: which <exec1> <exec2> <exec3> ...\n\n")
			cmd.Help()
		}
		which.Find(args...)
	},
}
