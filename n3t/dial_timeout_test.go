package n3t

import (
	"net"
	"testing"
	"time"
)

func TestDialTimeout(t *testing.T) {
	conn, err := DialTimeout("tcp", "10.7.7.7:http", 3*time.Second)
	if err == nil {
		conn.Close()
		t.Error("connection did not timed out!")
	}

	netErr, ok := err.(net.Error)
	if !ok {
		t.Fatal(err)
	}

	if !netErr.Timeout() {
		t.Log("Some other error, not timeout!")
		t.Fatal(err)
	}
}
