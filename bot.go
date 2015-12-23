package main

import (
	"encoding/json"
	"github.com/vampirewalk/slackbot/Godeps/_workspace/src/github.com/mvdan/xurls"
	"github.com/vampirewalk/slackbot/Godeps/_workspace/src/github.com/nlopes/slack"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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

func (b *Bot) ParseRepoFormat(repoURL string) (user string, repoName string) {
	u, err := url.Parse(repoURL)
	if err != nil {
		log.Printf("URL error")
	}
	results := strings.Split(u.Path, "/")
	return results[1], results[2]
}

func (b *Bot) GetRepoState(user string, repoName string) (repo Repository, err error) {
	resp, err := http.Get("https://api.github.com/repos/" + user + "/" + repoName)
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
	state := string("Language: " + repo.Language + "\nStar: " + strconv.Itoa(repo.StargazersCount) + "\nOpen issues: " + strconv.Itoa(repo.OpenIssuesCount) + "\n")
	return state
}
