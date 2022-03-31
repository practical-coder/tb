package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	printenvCmd.Flags().String("format", "", "Output format template")
}

var printenvCmd = &cobra.Command{
	Use:   "printenv",
	Short: "printenv utility",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		if format == "" {
			format = "{{.Key}}={{.Value}}\n"
		} else {
			format = fmt.Sprintf("%s\n", format)
		}
		printenv(args, format)
	},
}

type EnvVar struct {
	Key   string
	Value string
}

func printenv(args []string, format string) {
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
		envVars := make([]EnvVar, 0)

		for _, env := range os.Environ() {
			envSlice := strings.Split(env, "=")
			ev := EnvVar{
				Key:   envSlice[0],
				Value: envSlice[1],
			}
			envVars = append(envVars, ev)

			tmpl, err := template.New("env-vars").Parse(format)
			if err != nil {
				log.Logger.Fatal().Err(err).Msg("env-vars template error")
			}
			err = tmpl.Execute(os.Stdout, ev)
			if err != nil {
				log.Logger.Fatal().Err(err).Msg("env-vars template execute error")
			}
		}
	}
}
