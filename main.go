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
	os.Setenv("SLACK_BOT_TOKEN","xoxb-*****************************************************")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-***************************************************")
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
