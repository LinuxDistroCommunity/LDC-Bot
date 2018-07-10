package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config Config

//Config Struct
type Config struct {
	Irc struct {
		Server    string   `json:"Server"`
		Nick      string   `json:"Nick"`
		Ident     string   `json:"Ident"`
		TLS       bool     `json:"TLS"`
		SkipVerif bool     `json:"SkipVerif"`
		Channels  []string `json:"Channels"`
	} `json:"irc"`
}

//ReadConfig Unmarchal the Config file with Viper
func ReadConfig() Config {
	err := viper.Unmarshal(&config)
	if err != nil {
		log.WithFields(log.Fields{
			"Method": "Viper Config Unmarshal Error",
			"Error":  err,
		}).Fatal("ReadConfig()")
	}

	return config
}
