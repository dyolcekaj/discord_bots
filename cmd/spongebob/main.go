package main

import (
	"context"

	"github.com/dyolcekaj/discord_bots/internal/bot"
	"github.com/sirupsen/logrus"
)

func main() {
	bot, err := bot.New(context.Background(), bot.BotOptions{})

	if err != nil {
		logrus.Fatal(err)
	}

	if err := bot.Run(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("bot gracefully shutdown")
}
