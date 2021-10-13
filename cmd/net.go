package cmd

import (
	"github.com/practical-coder/tb/n3t"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

func init() {
	netCmd.AddCommand(
		netLookupCmd,
		netSrvCmd,
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
			log.Info().Msg("One argument required: <domain_name>")
			cmd.Help()
			return
		}
		n3t.LookupIP(args[0])
	},
}

var netSrvCmd = &cobra.Command{
	Use:     "srv",
	Example: "srv 7777",
	Short:   "TCP basic server listening on a given <port>",
	Long:    "TCP basic server listening on a given <port>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Info().Msg("One argument required: <port_number> range 1024-65535")
			cmd.Help()
			return
		}
		n3t.Listener(args[0])
	},
}
