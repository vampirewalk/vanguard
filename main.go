package main

import (
	"github.com/vampirewalk/vanguard/Godeps/_workspace/src/github.com/nlopes/slack"
	"log"
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
				if repoURL != "" && strings.Contains(repoURL, "github.com") && !bot.IsMe(ev.Msg.User) {
					log.Printf("Received message \"%s\" from %s", ev.Msg.Text, ev.Msg.User)
					user, repoName := bot.ParseRepoFormat(repoURL)
					state, err := bot.GetRepoState(user, repoName)
					if err != nil {
						log.Printf("Parse error")
					}
					report := bot.ReportRepoState(state)
					log.Printf(report)
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
