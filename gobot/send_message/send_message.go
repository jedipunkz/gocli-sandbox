package main

import (
	"log"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName(".gocli-sandbox")
	viper.AddConfigPath("$HOME/tmp")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Fatal error config file %s \n", err)
	}

	token := viper.GetString("token")

	api := slack.New(token)
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	channelID, timestamp, err := api.PostMessage("bot", slack.MsgOptionText("Some text", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Println("%s\n", err)
		return
	}
	log.Println("Message successfully sent to channel %s at %s", channelID, timestamp)
}
