package cmd

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Filename string

func init() {
	cryptoCmd.AddCommand(
		cryptoMD5Cmd,
		cryptoSHA256Cmd,
	)
	cryptoMD5Cmd.Flags().String("text", "", "string value")
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
		sValue, err := cmd.Flags().GetString("text")
		if err == nil && sValue != "" {
			fmt.Printf("%s\tMD5: %x\n", sValue, md5.Sum([]byte(sValue)))
			return
		}

		if len(args) == 0 {
			cmd.Help()
			return
		}

		for _, filename := range args {
			file, err := os.ReadFile(filename)
			if err != nil {
				log.Info().Err(err).Str("filename", filename).Msgf("Error reading file!")
				continue
			}

			fmt.Printf("%s\tMD5: %x\n", filename, md5.Sum(file))
		}
	},
}

var cryptoSHA256Cmd = &cobra.Command{
	Use:     "sha512",
	Example: "sha512 example_text",
	Short:   "sha512/256 hash function on files in arguments",
	Long:    "sha512/256 hash function on files in arguments",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		for _, filename := range args {
			file, err := os.ReadFile(filename)
			if err != nil {
				log.Info().Err(err).Str("filename", filename).Msgf("Error reading file!")
				continue
			}

			fmt.Printf("%s\tSHA512/256: %x\n", filename, sha512.Sum512_256(file))
		}
	},
}
