package functions

import (
	"fmt"

	irc "github.com/thoj/go-ircevent"
)

func Back(e *irc.Event) string {
	return fmt.Sprintf("Welcome back %s", e.Nick)
}
