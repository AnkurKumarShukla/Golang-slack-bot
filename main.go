package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
	
)

func printcommandevent(analyticsChannel <-chan *slacker.CommandEvent){
for event:= range analyticsChannel{
	fmt.Println("Command Events")
	fmt.Println(event.Timestamp)
	fmt.Println(event.Command)
	fmt.Println(event.Parameters)
	fmt.Println(event.Event)
	fmt.Println()
}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN","xoxb-2505861133808-6371763367030-NqWGFv1x5uzOgzBbBrCRY9oK")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A06AXM8QMKQ-6378254731635-f041ea8b3181e4f5ef081c7d7a85a6dfaf1f0972957264b8ce27feee9d625aa3")
bot:=slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))
go printcommandevent(bot.CommandEvents())
bot.Command("ping",&slacker.CommandDefinition{
	Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
		w.Reply("sun raha hu ! baihara nai hu mai")
	} ,
})

ctx,cancel := context.WithCancel(context.Background())
defer cancel()

err:=bot.Listen(ctx)
if err!= nil {
	log.Fatal(err)
}

}