package monitor

import "log"

type Monitor struct {
	*log.Logger
}

func (m *Monitor) Write(p []byte) (int, error) {
	err := m.Output(2, string(p))
	if err != nil {
		log.Println(err)
	}
	return len(p), nil
}
