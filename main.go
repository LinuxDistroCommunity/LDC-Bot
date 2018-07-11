package main

import (
	"github.com/LinuxDistroCommunity/LDC-Bot/config"
	"github.com/LinuxDistroCommunity/LDC-Bot/irc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	cnf config.Config
)

func init() {
	// Add the Viper Config set
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.AddConfigPath("/etc/LDC-Bot/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.LDC-Bot") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		logrus.WithFields(logrus.Fields{
			"Method": "Viper Config Load Error",
			"Error":  err,
		}).Fatal("Main init()")
	}
	cnf = config.ReadConfig()
}

func main() {
	// lets see if the bot actually starts lol
	irc := irc.IrcConnect(cnf)

	//All Code should be Implimented before this guy here this holds the main thread open
	irc.Join(cnf.Irc.Channels[0])
	irc.Loop()
}
