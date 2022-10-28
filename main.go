package main

import (
	"github.com/alexflint/go-arg"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	models2 "github.com/magmel48/github-releases-notifier/pkg/models"
	"os"
	"strings"
	"time"
)

func main() {
	_ = godotenv.Load()

	c := models2.Config{
		Interval: time.Hour,
		LogLevel: "info",
	}
	arg.MustParse(&c)

	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger,
		"ts", log.DefaultTimestampUTC,
		"caller", log.Caller(5),
	)

	// level.SetKey("severity")
	switch strings.ToLower(c.LogLevel) {
	case "debug":
		logger = level.NewFilter(logger, level.AllowDebug())
	case "warn":
		logger = level.NewFilter(logger, level.AllowWarn())
	case "error":
		logger = level.NewFilter(logger, level.AllowError())
	default:
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	if len(c.Repositories) == 0 {
		level.Error(logger).Log("msg", "no repositories to watch")
		os.Exit(1)
	}

	checker := &Checker{
		logger: logger,
		tokens: map[string]string{Github: c.GithubAuthToken, Gitlab: c.GitlabAuthToken},
	}

	releases := make(chan models2.Repository)
	go checker.Run(c.Interval, c.Repositories, releases)

	slack := SlackSender{Hook: c.Hook, Username: c.Username, Icon: c.Icon}

	level.Info(logger).Log("msg", "waiting for new releases")
	for repository := range releases {
		if c.IgnoreNonstable && repository.Release.IsNonstable() {
			level.Debug(logger).Log("msg", "not notifying about non-stable version", "version", repository.Release.Name)
			continue
		}

		if err := slack.Send(repository); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to send release to messenger",
				"err", err,
			)
			continue
		}
	}
}
