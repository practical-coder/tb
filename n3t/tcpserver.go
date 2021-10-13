package n3t

import (
	"net"
	"os"

	"github.com/rs/zerolog/log"
)

func Listener(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Error().
			Err(err).
			Msg("net Listen Error")
		os.Exit(1)
	}

	defer func() {
		_ = listener.Close()
	}()

	log.Info().
		Msgf("Bound to %q\n", listener.Addr())

}
