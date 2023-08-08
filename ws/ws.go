package ws

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/websocket"
)

func StartServer() {
	fs := http.FileServer(http.Dir("srv"))
	http.Handle("/", fs)
	http.Handle("/echo", websocket.Handler(Echo))
	http.Handle("/cputemp", websocket.Handler(Temperature))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("HTTP server error")
	}
}
