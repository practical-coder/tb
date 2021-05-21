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

var printenvCmd = &cobra.Command{
	Use:   "printenv",
	Short: "printenv utility",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, key := range args {
				fmt.Println(os.Getenv(key))
			}
		} else {
			for _, key := range os.Environ() {
				fmt.Println(key)
			}
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
