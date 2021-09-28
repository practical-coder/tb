package cmd

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"

	"github.com/spf13/cobra"
)

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
	Example: "md5 example_text",
	Short:   "md5 hash function on first argument",
	Long:    "md5 hash function on first argument",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}

		text := []byte(args[0])
		fmt.Printf("%x\n", md5.Sum(text))
	},
}

var cryptoSHA256Cmd = &cobra.Command{
	Use:     "sha256",
	Example: "sha256 example_text",
	Short:   "sha256 hash function on first argument",
	Long:    "sha256 hash function on first argument",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}

		text := []byte(args[0])
		fmt.Printf("%x\n", sha256.Sum256(text))
	},
}
