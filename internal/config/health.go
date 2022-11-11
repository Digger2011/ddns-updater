package config

import (
	"fmt"

	"github.com/qdm12/golibs/params"
)

type Health struct {
	ServerAddress string
	Port          uint16 // obtained from ServerAddress
}

func (h *Health) Get(env params.Interface) (warning string, err error) {
	h.ServerAddress, warning, err = env.ListeningAddress(
		"HEALTH_SERVER_ADDRESS", params.Default("127.0.0.1:9999"))
	if err != nil {
		return warning, fmt.Errorf("%w: for environment variable HEALTH_SERVER_ADDRESS", err)
	}
	h.Port = 8000
	return warning, nil
}
