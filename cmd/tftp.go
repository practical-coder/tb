package cmd

import (
	"os"

	"github.com/practical-coder/tb/tftp"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var tftpAddress string
var tftpFilename string

func init() {

	tftpCmd.AddCommand(tftpServerCmd)
	tftpServerCmd.Flags().StringVarP(&tftpAddress, "addr", "a", "0.0.0.0:9999", "UDP <host:port> address of the server")
	tftpServerCmd.Flags().StringVarP(&tftpFilename, "file", "f", "", "File path to the file that will be served")
}

var tftpCmd = &cobra.Command{
	Use:   "tftp",
	Short: "TFTP related tools",
	Long:  "TFTP related tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var tftpServerCmd = &cobra.Command{
	Use:     "server",
	Example: "server --addr <host:port> --file <file_path>",
	Short:   "UDP basic TFTP server listening on a given addr <host:port> serving file <file_path>",
	Long:    "UDP basic TFTP server listening on a given addr <host:port> serving file <file_path>",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := os.ReadFile(tftpFilename)
		if err != nil {
			log.Info().Err(err).Msgf("Problem with file in path: %s", tftpFilename)
		}

		s := tftp.Server{Payload: payload}
		log.Fatal().Err(s.ListenAndServe(tftpAddress)).Send()
	},
}
