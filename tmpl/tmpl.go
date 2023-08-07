package tmpl

import (
	"os"
	"text/template"

	"github.com/rs/zerolog/log"
)

func Apply(source string) {
	t, err := template.ParseFiles(source)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("source", source).
			Msg("ParseFiles error")
	}
	envs := map[string]string{
		"DB_URL": "postgres://",
	}
	err = t.Execute(os.Stdout, envs)
	if err != nil {
		log.Fatal().
			Err(err).
			Interface("envs", envs).
			Msg("Template Execute Error")
	}
}
