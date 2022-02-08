package txt

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
)

func Count(path string) {
	var err error
	var data []byte

	if path == "-" {
		data, err = ioutil.ReadAll(os.Stdin)
	} else {
		data, err = os.ReadFile(path)
	}

	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("Error reading file: %s", path)
	}

	log.Info().
		Str("file", path).
		Str("count", fmt.Sprintf("%d", len(data))).
		Msg("Number of bytes counted")

}
