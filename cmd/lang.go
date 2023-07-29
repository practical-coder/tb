package cmd

import (
	"fmt"

	"github.com/practical-coder/tb/lang"
	"github.com/practical-coder/tb/read"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	langCmd.AddCommand(
		langDetectCmd,
	)
}

var langCmd = &cobra.Command{
	Use:   "lang",
	Short: "Lang related tools",
	Long:  "Lang related tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var langDetectCmd = &cobra.Command{
	Use:     "detect",
	Example: "detect <file>",
	Short:   "Lang: detect source language using n-grams",
	Long:    "Lang: detect source language using n-grams",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Info().Msg("One argument required: <file_name>")
			cmd.Help()
			return
		}

		f, err := read.From(args[0])
		if err != nil {
			log.Fatal().Str("resource", args[0]).Msg("Cannot open file")
			return
		}
		defer f.Close()

		result, err := lang.Detect(f)
		if err != nil {
			log.Fatal().Msg("lang.Detect error")
			return
		}

		fmt.Println(result)
	},
}
