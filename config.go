package main

import (
	"time"
)

type SlackConfig struct {
	Hook        string			`arg:"env:SLACK_HOOK"`
	Username	string			`arg:"env:SLACK_USERNAME" default:"Releases Notifier"`
	Icon		string			`arg:"env:SLACK_ICON" default:"https://github.githubassets.com/favicons/favicon.png"`
}

type Config struct {
	GithubAuthToken string        	`arg:"env:GITHUB_AUTH_TOKEN"`
	GitlabAuthToken string        	`arg:"env:GITLAB_AUTH_TOKEN"`
	Interval        time.Duration 	`arg:"env:INTERVAL"`
	LogLevel        string        	`arg:"env:LOG_LEVEL"`
	Repositories    []string      	`arg:"-r,separate"`
	IgnoreNonstable bool          	`arg:"env:IGNORE_NONSTABLE"`
	SlackConfig
}
