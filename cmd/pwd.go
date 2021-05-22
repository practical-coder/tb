package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var pwdPhysical bool
var pwdLogical bool

func init() {
	pwdCmd.Flags().BoolVarP(&pwdPhysical, "physical", "P", false, "Physical working directory absolute path")
	pwdCmd.Flags().BoolVarP(&pwdLogical, "logical", "L", true, "Logical working directory absolute path")
}

var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Print Working Directory absolute path",
	Run: func(cmd *cobra.Command, args []string) {
		pwd(args)
	},
}

func pwd(args []string) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	} else {
		if !pwdPhysical {
			fmt.Println(path)
		} else {
			path, err = filepath.EvalSymlinks(path)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(path)
			}
		}
	}
}
