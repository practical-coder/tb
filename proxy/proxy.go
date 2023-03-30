package proxy

import (
	"io"
)

func proxy(source io.Reader, destination io.Writer) error {
	sourceWriter, isSourceWriter := source.(io.Writer)
	destinationReader, isDestinationReader := destination.(io.Reader)

	if isSourceWriter && isDestinationReader {
		go func() {
			_, _ = io.Copy(sourceWriter, destinationReader)
		}()
	}

	_, err := io.Copy(destination, source)

	return err
}

/*
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
	return err
}
*/
