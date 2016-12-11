package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/drgarcia1986/slacker/slack"
)

var token string
var channel string
var username string
var avatar string
var message string

func init() {
	flag.StringVar(&token, "t", os.Getenv("SLACKTOKEN"), "Slack integration token")
	flag.StringVar(&channel, "c", "", "Slack channel to post message")
	flag.StringVar(&username, "u", "Slacker", "BOT name")
	flag.StringVar(&avatar, "a", ":scream:", "BOT avatar")
	flag.StringVar(&message, "m", "", "Message to post")

	flag.Parse()
}

func main() {
	slackClient := slack.New(token)
	err := slackClient.PostMessage(channel, username, avatar, message)
	if err != nil {
		fmt.Printf("Cannot post message in slack, error: %s", err)
		os.Exit(2)
	}
}
