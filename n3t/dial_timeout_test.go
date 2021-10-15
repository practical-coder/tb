package n3t

import (
	"context"
	"net"
	"syscall"
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

func TestDialContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var dialer net.Dialer
	dialer.Control = func(_, _ string, _ syscall.RawConn) error {
		time.Sleep(3*time.Second + time.Millisecond)
		return nil
	}

	conn, err := dialer.DialContext(ctx, "tcp", "10.7.7.7:http")
	if err == nil {
		conn.Close()
		t.Fatal("Connection did not timed out")
	}

	netErr, ok := err.(net.Error)
	if !ok {
		t.Error(err)
	} else {
		if !netErr.Timeout() {
			t.Errorf("Error is not timeout: %v", err)
		}
	}

	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("DeadlineExceeded expected! Got: %v", ctx.Err())
	}
}

func TestDialContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var dialer net.Dialer
	dialer.Control = func(_, _ string, _ syscall.RawConn) error {
		t.Log("In custom dialer.Control!")
		time.Sleep(1 * time.Second)
		return nil
	}

	go func() {
		conn, err := dialer.DialContext(ctx, "tcp", "10.7.7.7:http")
		if err != nil {
			t.Error("DialContext failed")
		}
		conn.Close()
	}()
	time.Sleep(100 * time.Millisecond)
	cancel()

	if ctx.Err() != context.Canceled {
		t.Error(ctx.Err())
		t.Error("Something went wrong. Context should be Canceled")
	}
}
