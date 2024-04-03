package modbus

import (
	"errors"
	"net/url"
	"strings"

	"github.com/snple/kirara/util/mapform"
)

type Config struct {
	Addr            string `cfg:"addr"`
	Timeout         int    `cfg:"timeout,default=10"`
	Speed           int    `cfg:"speed,default=19200"`
	SlaveID         uint8  `cfg:"slaveId"`
	IsBIGEndian     bool   `cfg:"isBigEndian,default=true"`
	IsHighWordFirst bool   `cfg:"isHighWordFirst,default=true"`
	Debug           bool   `cfg:"debug"`
	isTCP           bool
	isRTU           bool
}

func ParseConfig(params string) (Config, error) {
	config := Config{
		Addr:            "tcp://127.0.0.1:502",
		Timeout:         10,
		Speed:           19200,
		SlaveID:         1,
		IsBIGEndian:     true,
		IsHighWordFirst: true,
		Debug:           false,
	}

	u, err := url.ParseQuery(params)
	if err != nil {
		return config, err
	}

	err = mapform.MapFormWithTag(&config, u, "cfg")
	if err != nil {
		return config, err
	}

	if strings.HasPrefix(config.Addr, "tcp://") {
		config.isTCP = true
	} else if strings.HasPrefix(config.Addr, "rtu://") {
		config.isRTU = true
	} else {
		return config, errors.New("config error, addr must begin with tcp or rtu")
	}

	return config, nil
}
