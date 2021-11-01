package monitor

import "log"

type Monitor struct {
	*log.Logger
}

func (m *Monitor) Write(p []byte) (int, error) {
	return len(p), m.Output(2, string(p))
}
