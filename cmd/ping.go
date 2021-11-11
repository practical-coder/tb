package cmd

import (
	"os"
	"time"

	"github.com/practical-coder/tb/n3t"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var pingCount int
var pingInterval time.Duration
var pingTimeout time.Duration

func init() {
	netPingCmd.Flags().IntVarP(&pingCount, "count", "c", 3, "Count of pings default: 3. If count <= 0 than infinite")
	netPingCmd.Flags().DurationVarP(&pingInterval, "interval", "i", time.Second, "Interval between pings. Default: 1 second")
	netPingCmd.Flags().DurationVarP(&pingTimeout, "timeout", "t", 5*time.Second, "Timeout, default 5 seconds")
}

var netPingCmd = &cobra.Command{
	Use:     "ping",
	Example: "ping --count=3 --interval=1s --timeout=5s google.com:80",
	Short:   "Ping TCP port of remote host",
	Long:    "Ping TCP port of remote host",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Info().Msg("One argument required: <domain|ip_addr>:<tcp_port>")
			cmd.Help()
			os.Exit(1)
		}

		params := n3t.PingParams{
			Count:    pingCount,
			Interval: pingInterval,
			Host:     args[0],
		}

		n3t.TcpPing(params)
	},
}
