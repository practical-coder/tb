package types

import (
	"net"
	"reflect"
	"testing"
)

func TestPayload(t *testing.T) {
	b1 := Binary("Clear is better than clever.")
	b2 := Binary("Don't Panic!")
	s1 := String("Errors are values.")

	payloads := []Payload{&b1, &b2, &s1}

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		for _, payload := range payloads {
			_, err := payload.WriteTo(conn)
			if err != nil {
				t.Error(err)
				break
			}
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < len(payloads); i++ {
		actual, err := decode(conn)
		if err != nil {
			t.Fatal(err)
		}

		if expected := payloads[i]; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Not equal! %v != %v", actual, expected)
		}

		t.Logf("[%T] %[1]q", actual)
	}
}
