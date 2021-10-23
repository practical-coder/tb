package n3t

import (
	"bufio"
	"net"
	"reflect"
	"testing"
)

const payload = "Sharpen the saw! Do not ignore fundamentals!"

func TestScanner(t *testing.T) {
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

		_, err = conn.Write([]byte(payload))
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		t.Error(err)
	}

	expected := []string{
		"Sharpen", "the", "saw!", "Do", "not", "ignore", "fundamentals!",
	}

	if !reflect.DeepEqual(expected, words) {
		t.Fatalf("Scanning words went wrong! Expected: %v, Actual: %v", expected, words)
	}

	t.Logf("Scanned words: %v", words)
}
