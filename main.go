package main

import (
	"github.com/vampirewalk/slackbot/Godeps/_workspace/src/github.com/nlopes/slack"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("SLACK_TOKEN is required")
	}

	api := slack.New(token)
	bot := Bot{}

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				bot.SetMe(ev.Info.User)

			case *slack.MessageEvent:
				// if msg contains github url
				repoURL := bot.ExtractGithubRepoURL(ev.Msg.Text)
				if repoURL != "" && !bot.IsMe(ev.Msg.User) {
					log.Printf("Received message \"%s\" from %s", ev.Msg.Text, ev.Msg.User)
					u, err := url.Parse(repoURL)
					if err != nil {
						rtm.SendMessage(rtm.NewOutgoingMessage("url error", ev.Channel))
					}
					results := strings.Split(u.Path, "/")
					state, err := bot.GetRepoState(results[2])
					if err != nil {
						rtm.SendMessage(rtm.NewOutgoingMessage("Parse error", ev.Channel))
					}
					report := bot.ReportRepoState(state)
					rtm.SendMessage(rtm.NewOutgoingMessage(report, ev.Channel))
				}

			case *slack.RTMError:
				log.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				log.Fatal("Invalid credentials")
				break Loop
			}
		}
	}
}
