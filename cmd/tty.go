package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ttyCmd represents the tty command
var ttyCmd = &cobra.Command{
	Use:   "tty",
	Short: "TTY",
	Long:  `TTY`,
	Run: func(cmd *cobra.Command, args []string) {
		fileInfo, _ := os.Stdout.Stat()
		if (fileInfo.Mode() & os.ModeCharDevice) != 0 {
			fmt.Println("TTY: interactive terminal")
		} else {
			fmt.Println("Not a TTY terminal")
		}
	},
}

func init() {
	rootCmd.AddCommand(ttyCmd)
}
