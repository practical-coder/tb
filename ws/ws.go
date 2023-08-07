package ws

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"golang.org/x/net/websocket"
)

func EchoServer() {
	http.Handle("/", websocket.Handler(Echo))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("HTTP server error")
	}
}

func Echo(conn *websocket.Conn) {
	fmt.Println("Echoing")
	for n := 0; n < 10; n++ {
		msg := fmt.Sprintf("Hello %c\n", n+48)
		fmt.Println("Sending to client: " + msg)
		err := websocket.Message.Send(conn, msg)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("websocket message send error")
			break
		}
		var reply string
		err = websocket.Message.Receive(conn, &reply)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("websocket message receive error")
			break
		}

		fmt.Println("Received back from client: " + reply)
	}
}

func EchoClient(source string) {
	conn, err := websocket.Dial(source, "", "http://localhost:12345")
	if err != nil {
		log.Fatal().
			Err(err).
			Str("source", source).
			Msg("websocket Dial error")
	}
	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Info().
				Err(err).
				Msg("websocket message receive error")
			break
		}
		fmt.Println("Received from server:", msg)
		err = websocket.Message.Send(conn, msg)
		if err != nil {
			log.Info().
				Err(err).
				Msg("websocket message send error")
			break
		}
	}
}
