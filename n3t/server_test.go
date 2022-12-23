package n3t

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestSimpleHTTPServer(t *testing.T) {
	srv := http.Server{
		Addr:              "127.0.0.1:8081",
		Handler:           http.TimeoutHandler(DefaultHandler{}, 20*time.Second, ""),
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	listener, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		err := srv.Serve(listener)
		if err != http.ErrServerClosed {
			t.Error(err)
		}
	}()

	testCases := []struct {
		method   string
		body     io.Reader
		code     int
		response string
	}{
		{http.MethodGet, nil, http.StatusOK, "It works!"},
		{http.MethodPost, strings.NewReader("Uploaded"), http.StatusOK, "Uploaded"},
		{http.MethodHead, nil, http.StatusMethodNotAllowed, ""},
	}

	client := new(http.Client)
	path := fmt.Sprintf("http://%s/", srv.Addr)

	for i, c := range testCases {
		req, err := http.NewRequest(c.method, path, c.body)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}

		if resp.StatusCode != c.code {
			t.Errorf("%d: expected status code: %d; actual status code: %d", i, c.code, resp.StatusCode)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}

		_ = resp.Body.Close()

		if c.response != string(b) {
			t.Errorf("%d: expected: %q; actual: %q", i, c.response, b)
		}
	}

	if err := srv.Close(); err != nil {
		t.Fatal(err)
	}

}

type DefaultHandler struct{}

func (dh DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("It works!"))
	case http.MethodHead:
		w.WriteHeader(http.StatusMethodNotAllowed)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write([]byte(body))
	}

}
