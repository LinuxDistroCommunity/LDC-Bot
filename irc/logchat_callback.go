package irc

import (
	"github.com/sirupsen/logrus"
	"github.com/thoj/go-ircevent"
)

//LogIrcToConsole Log IRC Output to console
func LogIrcToConsole(ircobj *irc.Connection) {
	ircobj.AddCallback("PRIVMSG", func(event *irc.Event) {
		go func(event *irc.Event) {
			//event.Message() contains the message
			//event.Nick Contains the sender
			//event.Arguments[0] Contains the channel
			logrus.WithFields(logrus.Fields{
				"Channel": event.Arguments[0],
				"Nick":    event.Nick,
				"Message": event.Message(),
			}).Info()
		}(event)
	})
}
