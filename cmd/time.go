package cmd

import (
	"github.com/practical-coder/tb/datetime"

	"github.com/spf13/cobra"
)

func init() {
	timeCmd.AddCommand(
		timeNowCmd,
		timeEpochCmd,
		timeFormatCmd,
	)
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Using time package from standard library",
	Long:  "Using time package from standard library",
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()

	},
}

var timeFormatCmd = &cobra.Command{
	Use:   "format",
	Short: "list available formats from time package",
	Long:  "list available formats from time package",
	Run: func(cmd *cobra.Command, args []string) {
		datetime.Format()
	},
}

var timeNowCmd = &cobra.Command{
	Use:   "now",
	Short: "current time",
	Long:  "current time",
	Run: func(cmd *cobra.Command, args []string) {
		datetime.TheTimeIsNow()
	},
}

var timeEpochCmd = &cobra.Command{
	Use:   "epoch",
	Short: "Unix Epoch time formats",
	Long:  "Unix Epoch time formats",
	Run: func(cmd *cobra.Command, args []string) {
		datetime.Epoch()
	},
}
