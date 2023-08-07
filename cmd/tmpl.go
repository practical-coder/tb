package cmd

import (
	"github.com/practical-coder/tb/tmpl"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	tmplCmd.AddCommand(applyTmplCmd)
	applyTmplCmd.Flags().String("source", "", "template path or URL")
}

var tmplCmd = &cobra.Command{
	Use:   "tmpl",
	Short: "template tools",
	Long:  "template tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var applyTmplCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply given template",
	Long:  "Apply given template",
	Run: func(cmd *cobra.Command, args []string) {
		source, err := cmd.Flags().GetString("source")
		if err != nil {
			log.Fatal().Err(err).Str("source", source).Msg("template source is required")
			return
		}
		tmpl.Apply(source)
	},
}
