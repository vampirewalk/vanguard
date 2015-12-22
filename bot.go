package main

import (
	"encoding/json"
	"github.com/vampirewalk/slackbot/Godeps/_workspace/src/github.com/mvdan/xurls"
	"github.com/vampirewalk/slackbot/Godeps/_workspace/src/github.com/nlopes/slack"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Bot struct {
	MeID string
}

func (b *Bot) SetMe(user *slack.UserDetails) {
	b.MeID = user.ID
	log.Printf("Connect! I am %s (%s)\n", user.Name, user.ID)
}

func (b *Bot) IsMe(id string) bool {
	return b.MeID == id
}

func (b *Bot) IsDM(channel string) bool {
	return regexp.MustCompile("^D.*").MatchString(channel)
}

func (b *Bot) AmIMentioned(text string) bool {
	return regexp.MustCompile("<@" + b.MeID + ">").MatchString(text)
}

func (b *Bot) ExtractGithubRepoURL(text string) string {
	return xurls.Strict.FindString(text)
}

func (b *Bot) GetRepoState(repoName string) (repo Repository, err error) {
	resp, err := http.Get("https://api.github.com/repos/" + repoName)
	if err != nil {
		// handle error
		return Repository{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return Repository{}, err
	}

	err = json.Unmarshal(body, &repo)
	if err != nil {
		return Repository{}, err
	}

	return repo, nil
}

func (b *Bot) ReportRepoState(repo Repository) string {
	state := string("language: " + repo.Language + "\nstar: " + string(repo.StargazersCount) + "\nopen issues: " + string(repo.OpenIssuesCount) + "\n")
	return state
}
