package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		netCmd,
		printenvCmd,
		pwdCmd,
		timeCmd,
		whichCmd,
		versionCmd,
	)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

var rootCmd = &cobra.Command{
	Use:   "tb",
	Short: "ToolBox useful command line tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "ToolBox version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ToolBox version 0.0.3")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Err(err)
		os.Exit(1)
	}
}
