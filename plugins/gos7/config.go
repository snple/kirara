package gos7

import (
	"net/url"

	"github.com/snple/kirara/util/mapform"
)

type Config struct {
	Addr        string `cfg:"addr"`
	Rank        int    `cfg:"rank"`
	Slot        int    `cfg:"slot"`
	IsBIGEndian bool   `cfg:"isBigEndian,default=true"`
	Debug       bool   `cfg:"debug"`
}

func ParseConfig(params string) (Config, error) {
	config := Config{
		Addr:        "127.0.0.1",
		Rank:        0,
		Slot:        1,
		IsBIGEndian: true,
		Debug:       false,
	}

	u, err := url.ParseQuery(params)
	if err != nil {
		return config, err
	}

	err = mapform.MapFormWithTag(&config, u, "cfg")
	if err != nil {
		return config, err
	}

	return config, nil
}
