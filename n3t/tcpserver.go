package n3t

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
)

func Listener(port string) {
	address := fmt.Sprintf("127.0.0.1:%s", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Error().
			Err(err).
			Msg("net Listen Error")
	}

	defer func() {
		_ = listener.Close()
	}()

	log.Info().
		Msgf("Bound to %q\n", listener.Addr())

}
