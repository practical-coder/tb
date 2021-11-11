package n3t

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

type PingParams struct {
	Count    int
	Interval time.Duration
	Timeout  time.Duration
	Host     string
}

func TcpPing(params PingParams) {
	log.Info().Msgf("TCP PING: %s", params.Host)

	if params.Count <= 0 {
		log.Info().Msg("Press CMD/CTRL + C")
	}

	count := 0
	for params.Count <= 0 || count < params.Count {
		count++
		fmt.Printf("%d ", count)

		start := time.Now()
		conn, err := net.DialTimeout("tcp", params.Host, params.Timeout)
		duration := time.Since(start)

		if err != nil {
			log.Info().
				Err(err).
				Str("host", params.Host).
				Time("start", start).
				Dur("duration", duration).
				Interface("connection", conn).
				Msg("DialTimeout Error")
			netErr, isNetErr := err.(net.Error)
			if !isNetErr && !netErr.Temporary() {
				os.Exit(2)
			}
		} else {
			_ = conn.Close()
			log.Info().
				Str("host", params.Host).
				Dur("duration", duration).
				Send()
		}

		time.Sleep(params.Interval)
	}

}
