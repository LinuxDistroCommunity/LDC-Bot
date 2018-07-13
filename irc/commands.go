package irc

import (
	"fmt"

	"github.com/LinuxDistroCommunity/LDC-Bot/config"
	"github.com/thoj/go-ircevent"
)

func IrcCommands(cmd string, e *irc.Event, cnf config.Config) string {
	switch c := cmd; c {
	case fmt.Sprintf("%shello", cnf.Irc.Prefix):
		return fmt.Sprintf("Well Hello There %s", e.Nick)
	default:
		return fmt.Sprintf("%s Command Not Found", c)
	}
}
