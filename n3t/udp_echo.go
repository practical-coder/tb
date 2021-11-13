package n3t

import (
	"context"
	"net"

	"github.com/rs/zerolog/log"
)

func EchoServerUDP(ctx context.Context, addr string) (net.Addr, error) {
	pconn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Info().
			Err(err).
			Str("address", addr).
			Msg("ListenPacket error")
		return nil, err
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = pconn.Close()
		}()

		buf := make([]byte, 1024)
		for {
			n, clientAddr, err := pconn.ReadFrom(buf) // client to server
			if err != nil {
				log.Info().
					Err(err).
					Interface("clientAddr", clientAddr).
					Msg("ReadFrom error")
				return
			}

			_, err = pconn.WriteTo(buf[:n], clientAddr) // server to client
			if err != nil {
				log.Info().
					Err(err).
					Interface("clientAddr", clientAddr).
					Msg("WriteTo error")
				return
			}
		}
	}()

	return pconn.LocalAddr(), nil
}
