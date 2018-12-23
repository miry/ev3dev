// findpath is a reimplementation of the Demo program loaded on new ev3 bricks,
// without sound. It demonstrates the use of the ev3dev Go API.
// The control does not make full use of the ev3dev API where it could.

package main

import (
	"os"

	"github.com/miry/ev3dev/pkg/bot"
)

func main() {
	bot := bot.New()
	defer bot.Exit()

	bot.Run()
	os.Exit(0)
}
