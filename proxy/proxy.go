package proxy

import (
	"io"
	"net"
)

func proxy(source io.Reader, destination io.Writer) error {
	sourceReader, isSourceReader := source.(io.Reader)
	destinationWriter, isDestinationWriter := destination.(io.Writer)

	if isSourceReader && isDestinationWriter {
		go func() {
			_, _ = io.Copy(destinationWriter, sourceReader)
		}()
	}

	_, err := io.Copy(destination, source)

	return err
}

func proxyConn(source, destination string) error {
	connSource, err := net.Dial("tcp", source)
	if err != nil {
		return err
	}
	defer connSource.Close()

	connDestination, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer connDestination.Close()

	go func() {
		_, _ = io.Copy(connSource, connDestination)
	}()

	_, err = io.Copy(connDestination, connSource)
	return nil
}
