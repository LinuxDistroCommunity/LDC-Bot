package functions

import (
	"fmt"

	irc "github.com/thoj/go-ircevent"
)

func Hello(e *irc.Event) string {
	return fmt.Sprintf("Well Hello There %s", e.Nick)
}
