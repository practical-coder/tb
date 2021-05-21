package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var printenvCmd = &cobra.Command{
	Use:   "printenv",
	Short: "printenv utility",
	Run: func(cmd *cobra.Command, args []string) {
		printenv(args)
	},
}

func printenv(args []string) {
	if len(args) > 0 {
		for _, key := range args {
			val, ok := os.LookupEnv(key)
			if ok {
				fmt.Println(val)
			} else {
				fmt.Println(key, "not set!")
			}

		}
	} else {
		for _, key := range os.Environ() {
			fmt.Println(key)
		}
	}
}
