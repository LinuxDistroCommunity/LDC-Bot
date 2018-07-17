package irc

import (
	"testing"

	"github.com/LinuxDistroCommunity/LDC-Bot/config"
	"github.com/thoj/go-ircevent"
)

func TestIrcCommandCallback(t *testing.T) {
	type args struct {
		ircobj *irc.Connection
		cnf    config.Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IrcCommandCallback(tt.args.ircobj, tt.args.cnf)
		})
	}
}
