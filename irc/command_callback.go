package irc

import (
	"strings"

	"github.com/LinuxDistroCommunity/LDC-Bot/config"
	"github.com/thoj/go-ircevent"
)

// Command Here we are creating the Command struct that holds our message string
type Command struct {
	// here we set the struct value
	Message string
}

// IrcCommandCallback takes the input of the IRC channel and splits the string to find if there is any commands in the message
func IrcCommandCallback(ircobj *irc.Connection, cnf config.Config) {
	ircobj.AddCallback("PRIVMSG", func(event *irc.Event) {
		go func(event *irc.Event) {
			//event.Message() contains the message
			//event.Nick Contains the sender
			//event.Arguments[0] Contains the channel
			messages := make(chan Command)

			msg := strings.Split(event.Message(), " ")

			if strings.HasPrefix(msg[0], cnf.Irc.Prefix) {
				go func() {
					messages <- Command{IrcCommands(msg[0], event, cnf)}
				}()

				m := <-messages

				ircobj.Privmsg(event.Arguments[0], m.Message)
			}
			close(messages)
		}(event)
	})
}
