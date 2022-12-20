package tftp

import (
	"bytes"
	"errors"
	"net"
	"time"

	"github.com/rs/zerolog/log"
)

type Server struct {
	Payload []byte        // Payload served for all read requests
	Retries uint8         // Number of times to retry a failed transmission
	Timeout time.Duration // the duration to wait for an acknowledgement
}

func (s Server) ListenAndServe(addr string) error {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()

	log.Info().Msgf("Listening on %s...\n", conn.LocalAddr())

	return s.Serve(conn)
}

func (s *Server) Serve(conn net.PacketConn) error {
	if conn == nil {
		return errors.New("nil connection!")
	}

	if s.Payload == nil {
		return errors.New("Payload is required!")
	}

	if s.Retries == 0 {
		s.Retries = 5 // default 5 retries
	}

	if s.Timeout == 0 {
		s.Timeout = 5 * time.Second // default 5 seconds timeout
	}

	var rrq ReadReq

	for {
		buf := make([]byte, DatagramSize)

		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			return err
		}

		err = rrq.UnmarshalBinary(buf)
		if err != nil {
			log.Info().Err(err).Msgf("[%s] Bad Request!", rrq.Filename)
			continue
		}

		go s.handle(addr.String(), rrq)
	}
}

func (s *Server) handle(clientAddr string, rrq ReadReq) {
	log.Info().Msgf("[%s] client requested file: %s", clientAddr, rrq.Filename)

	conn, err := net.Dial("udp", clientAddr)
	if err != nil {
		log.Info().Err(err).Msg("[%s] client connection error!")
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	var (
		ackPacket  Ack
		errPacket  Err
		dataPacket = Data{Payload: bytes.NewReader(s.Payload)}
		buf        = make([]byte, DatagramSize)
	)

NEXTPACKET:
	for n := DatagramSize; n == DatagramSize; {
		data, err := dataPacket.MarshalBinary()
		if err != nil {
			log.Info().Err(err).Msgf("[%s] preparing data packet error", clientAddr)
			return
		}

	RETRY:
		for i := s.Retries; i > 0; i-- {
			n, err = conn.Write(data) // send data packet
			if err != nil {
				log.Info().Err(err).Msgf("[%s] write data packet error!", clientAddr)
				return
			}

			// wait for client's ACK packet
			_ = conn.SetReadDeadline(time.Now().Add(s.Timeout))
			_, err = conn.Read(buf)

			if err != nil {
				if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
					continue RETRY
				}

				log.Info().Err(err).Msgf("[%s] waiting for ACK", clientAddr)
				return
			}

			switch {
			case ackPacket.UnmarshalBinary(buf) == nil:
				if uint16(ackPacket) == dataPacket.Block {
					// received ACK, send NEXT PACKET
					continue NEXTPACKET
				}

			case errPacket.UnmarshalBinary(buf) == nil:
				log.Info().Msgf("[%s] received error: %v", clientAddr, errPacket.Message)
				return
			default:
				log.Info().Msgf("[%s] invalid packet!", clientAddr)
			}
		}

		log.Info().Msgf("[%s] %d retries exceeded!", clientAddr, s.Retries)
		return
	}

	log.Info().Msgf("[%s] sent %d blocks", clientAddr, dataPacket.Block)
}
