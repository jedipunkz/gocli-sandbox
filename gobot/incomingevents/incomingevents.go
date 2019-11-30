package main

import (
    "log"
    "os"
	"strings"

    "github.com/nlopes/slack"
	"github.com/spf13/viper"
)

func run(api *slack.Client) int {
    rtm := api.NewRTM()
    go rtm.ManageConnection()

    for {
		msg := <-rtm.IncomingEvents
		log.Printf("MSG: %#v\n", msg.Data)

		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			log.Printf("Start up!")

		case *slack.MessageEvent:
			if strings.Contains(ev.Text, "Foo") {
				text := "Foo Fighters!"
				rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
			}
		}


        // select {
        // case msg := <-rtm.IncomingEvents:
        //     switch ev := msg.Data.(type) {
        //     case *slack.HelloEvent:
        //         log.Print("Hello Event")
        //
        //     case *slack.MessageEvent:
        //         log.Printf("Message: %v\n", ev)
        //         rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))
        //
        //     case *slack.InvalidAuthEvent:
        //         log.Print("Invalid credentials")
        //         return 1
        //
        //     }
        // }
    }
}

func main() {
	viper.SetConfigName(".gocli-sandbox")
	viper.AddConfigPath("$HOME/tmp")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Fatal error config file %s \n", err)
	}

	token := viper.GetString("token")

    api := slack.New(token)
    os.Exit(run(api))
}
