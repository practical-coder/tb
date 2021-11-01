package monitor

import (
	"io"
	"log"
	"net"
	"os"
	"testing"
)

func TestMonitor(t *testing.T) {
	monitor := &Monitor{
		Logger: log.New(os.Stdout, "monitor: ", 0),
	}

	// Server
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		monitor.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		conn, err := listener.Accept()
		if err != nil {
			return
		}

		// TeeReader
		b := make([]byte, 1024)
		r := io.TeeReader(conn, monitor)
		n, err := r.Read(b)
		if err != nil && err != io.EOF {
			monitor.Println(err)
			return
		}

		// MultiWriter
		w := io.MultiWriter(conn, monitor)
		_, err = w.Write(b[:n])
		if err != nil && err != io.EOF {
			monitor.Println(err)
			return
		}

	}()

	// Client
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		monitor.Fatal(err)
	}

	_, err = conn.Write([]byte("It works!\n"))
	if err != nil {
		monitor.Fatal(err)
	}

	_ = conn.Close()
	<-done
}
