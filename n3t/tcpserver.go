package n3t

import (
	"fmt"
	"net"
	"os"
	"time"

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

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error().
				Err(err).
				Msg("net Accept Error")
		}

		go func(conn net.Conn) {
			defer conn.Close()

			ts := fmt.Sprintf("\n%s\n", time.Now().Format(time.RFC3339))
			conn.Write([]byte(ts))
			clientLog := log.Output(conn)
			clientLog.Info().
				Str("local_address", conn.RemoteAddr().String()).
				Str("remote_address", conn.LocalAddr().String()).
				Msg("Connection Established")
			log.Info().
				Interface("local_address", conn.LocalAddr()).
				Interface("remote_address", conn.RemoteAddr()).
				Msg("Connection Established")
		}(conn)
	}

}
