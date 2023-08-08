package ws

import (
	"os/exec"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/websocket"
)

func Temperature(conn *websocket.Conn) {
	for {
		cputemp, err := exec.Command("osx-cpu-temp").CombinedOutput()
		if err != nil {
			log.Info().
				Err(err).
				Msg("osx-cpu-temp command error")
			break
		}
		log.Info().Bytes("cputemp", cputemp).Msg("Sending CPU temperature")
		err = websocket.Message.Send(conn, string(cputemp))
		if err != nil {
			log.Info().
				Err(err).
				Msg("websockets message send error")
			break
		}
		time.Sleep(2 * time.Second)
		var reply string
		err = websocket.Message.Receive(conn, &reply)
		if err != nil {
			log.Info().
				Err(err).
				Msg("websockets message receive error")
			break
		}
		log.Info().Str("reply", reply).Msg("Received from Client")
	}

}
