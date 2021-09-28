package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		printenvCmd,
		pwdCmd,
		whichCmd,
		versionCmd,
	)
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
		fmt.Println("ToolBox version 0.0.1")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
