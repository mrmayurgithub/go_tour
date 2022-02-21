package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3107445446615-3125492597476-NGJEN4FgRbdJ5QN7BFORZ4iK")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A033PDGUV5J-3116493368374-e58f580fb6d0340894441ff6f909e3b9cb7e4a2399714bc1d3ec73da290c7864")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())
	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "YOB Calc",
		Example:     "My YOB is 2001",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2022 - yob
			r := fmt.Sprintf("Age is %d", age)
			response.Reply(r)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
