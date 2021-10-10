package cmd

import (
	"github.com/practical-coder/tb/net"

	"github.com/spf13/cobra"
)

func init() {
	netCmd.AddCommand(
		netLookupCmd,
	)
}

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Network related tools",
	Long:  "Network related tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var netLookupCmd = &cobra.Command{
	Use:     "lookup",
	Example: "lookup google.com",
	Short:   "Lookup domain IP address.",
	Long:    "Lookup domain IP address.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}
		net.LookupIP(args[0])
	},
}
