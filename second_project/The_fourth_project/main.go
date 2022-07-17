package main

import (
	"os"
)

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"github.com/shomalil/slacker"
// )

func PrintCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {

}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-310246501329-2253154075362-5bewxhgvdaoCpRKivBSeG4P4")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A027MAS7CC-2246395933878-354ffc9a0aad7282fca446b1999989f75a0b8d48e5284fedid86e59e18614118")

	// create bot
	bot := Slacker.NewCilent(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go PrintCommandEvents(bot.CommandEvents())
}
