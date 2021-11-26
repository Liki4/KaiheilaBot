package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type config struct {
	KhlBot struct {
		Token string `toml:"token"`
	} `toml:"khlbot"`

	Ncm struct {
		NcmApi  string `toml:"ncmapi"`
		Phone   string `toml:"phone"`
		Md5Pass string `toml:"md5pass"`
	} `toml:"ncm"`
}

var conf *config

func Load() error {
	c := config{}

	_, err := toml.DecodeFile("./config/khlbot.toml", &c)
	if err != nil {
		return errors.Wrap(err, "decode config file")
	}

	conf = &c
	return nil
}

// Get returns the config struct.
func Get() *config {
	return conf
}
