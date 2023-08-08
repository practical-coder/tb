package cmd

import (
	"github.com/practical-coder/tb/ws"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	wsCmd.AddCommand(echoServerCmd, echoClientCmd)
	echoClientCmd.Flags().String("source", "", "websockets service url")
}

var wsCmd = &cobra.Command{
	Use:   "ws",
	Short: "Websockets tools",
	Long:  "Websockets tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var echoServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Websockets server",
	Long:  "Websockets server",
	Run: func(cmd *cobra.Command, args []string) {
		ws.StartServer()
	},
}

var echoClientCmd = &cobra.Command{
	Use:   "echoclient",
	Short: "Websockets echo client",
	Long:  "Websockets echo client",
	Run: func(cmd *cobra.Command, args []string) {
		source, err := cmd.Flags().GetString("source")
		if err != nil {
			log.Fatal().Err(err).Str("source", source).Msg("websocket service source URL is required")
			return
		}
		ws.EchoClient(source)
	},
}
