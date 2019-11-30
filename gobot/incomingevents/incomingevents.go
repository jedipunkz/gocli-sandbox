package main

import (
	"log"
	"os"
	"os/exec"
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
			if strings.Contains(ev.Text, "terraform") {
				out, err := exec.Command("terraform", "version").Output()
				if err != nil {
					log.Println("Fatal error : %s \n", err)
				}

				rtm.SendMessage(rtm.NewOutgoingMessage("```"+string(out)+"```", ev.Channel))
			}
		}
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
