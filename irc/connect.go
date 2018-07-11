package irc

import (
	"crypto/tls"

	"github.com/LinuxDistroCommunity/LDC-Bot/config"
	"github.com/sirupsen/logrus"
	"github.com/thoj/go-ircevent"
)

//IrcConnect Connect to the IRC Network
func IrcConnect(cnf config.Config) *irc.Connection {
	irccon := irc.IRC(cnf.Irc.Nick, cnf.Irc.Ident)
	irccon.UseTLS = cnf.Irc.TLS
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: cnf.Irc.SkipVerif}

	/*
	 Add Callbacks here e.g

	 irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	 irccon.AddCallback("366", func(e *irc.Event) { })
	*/

	LogIrcToConsole(irccon)

	err := irccon.Connect(cnf.Irc.Server)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Method": "IRC Connection Error",
			"Error":  err,
		}).Fatal("IrcConnect()")
	}

	return irccon
}
