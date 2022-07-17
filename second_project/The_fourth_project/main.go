package main

import (
	"context"
	"fmt"
	"log"
	"os"
    "strconv"
	"github.com/shomalil/slacker"
)

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"github.com/shomalil/slacker"
// )

func PrintCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("command events")
		fmt.Println(event.TimeStamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-310246501329-2253154075362-5bewxhgvdaoCpRKivBSeG4P4")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A027MAS7CC-2246395933878-354ffc9a0aad7282fca446b1999989f75a0b8d48e5284fedid86e59e18614118")

	// create bot
	bot := Slacker.NewCilent(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go PrintCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>, &slacker.CommandDefinition"){
		// when user says my yob is <year> it is param
		Description: "yob calculator",
		Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)
		    year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2022-yob
			r := fmt.Sprint("age is %d", age)
			response.Reply(r)
	}

	ctx, cancel := context.WithCancel(ctx)
	// ctx related to context
	defer cancel()

	// Why do we use defer?
	// We usually use defer to close or deallocate resources.
	// A surrounding function executes all deferred function calls before it returns, even if it panics.
	// If you just place a function call at the end of a surrounding function, it is skipped when panic happens.
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// extra explanation is required