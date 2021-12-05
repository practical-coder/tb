package cmd

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Filename string

func init() {
	cryptoCmd.AddCommand(
		cryptoMD5Cmd,
		cryptoSHA256Cmd,
	)
}

var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "crypto package utilities",
	Long:  "crypto package utilities",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var cryptoMD5Cmd = &cobra.Command{
	Use:     "md5",
	Example: "md5 file_path ...",
	Short:   "md5 hash function on files in arguments",
	Long:    "md5 hash function on files in arguments",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		for _, filename := range args {
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Info().Err(err).Str("filename", filename).Msgf("Error reading file!")
				continue
			}

			fmt.Printf("%s\tMD5: %x\n", filename, md5.Sum(file))
		}
	},
}

var cryptoSHA256Cmd = &cobra.Command{
	Use:     "sha256",
	Example: "sha256 example_text",
	Short:   "sha256 hash function on files in arguments",
	Long:    "sha256 hash function on files in arguments",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		for _, filename := range args {
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Info().Err(err).Str("filename", filename).Msgf("Error reading file!")
				continue
			}

			fmt.Printf("%s\tSHA256: %x\n", filename, sha256.Sum256(file))
		}
	},
}
