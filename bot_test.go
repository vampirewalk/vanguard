package main

import (
	"net/url"
	"strings"
	"testing"
)

func TestExtractURL(t *testing.T) {
	s := "pousdnfsadlmgf\n<https://github.com/nlopes/slack>"
	bot := Bot{}
	url := bot.ExtractGithubRepoURL(s)
	expectedURL := "https://github.com/nlopes/slack"
	if url != expectedURL {
		t.Errorf("Failed to extract url")
	}
}

func TestGetRepoName(t *testing.T) {
	repoURL := "https://github.com/nlopes/slack"
	u, err := url.Parse(repoURL)
	if err != nil {
		t.Errorf("Failed to parse url")
	}
	results := strings.Split(u.Path, "/")
	if results[2] != "slack" {
		t.Errorf("repo name error")
	}
}

func TestGetRepoState(t *testing.T) {
	bot := Bot{}
	state, err := bot.GetRepoState("nlopes", "slack")
	if err != nil {
		t.Errorf("Failed to get repo state")
	}
	if state.Language == "" {
		t.Errorf("Empty state")
	}
}
