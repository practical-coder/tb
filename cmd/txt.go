package cmd

import (
	"github.com/practical-coder/tb/txt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	txtCmd.AddCommand(
		txtCountCmd,
	)
}

var txtCmd = &cobra.Command{
	Use:   "txt",
	Short: "Text related tools",
	Long:  "Text related tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var txtCountCmd = &cobra.Command{
	Use:     "count",
	Example: "count <file>",
	Short:   "Text: Count number of characters, words, lines",
	Long:    "Text: Count number of characters, words, lines",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Info().Msg("One argument required: <file_name>")
			cmd.Help()
			return
		}
		txt.Count(args[0])
	},
}
