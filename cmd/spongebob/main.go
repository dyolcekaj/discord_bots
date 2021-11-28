package main

import (
	"context"

	"github.com/dyolcekaj/discord_bots/pkg/discord"
	"github.com/sirupsen/logrus"
)

func main() {
	spongebob := discord.NewSlashCommand("spongebob")

	bot, err := discord.NewBot(
		context.Background(),
		discord.BotOptions{},
		[]discord.Command{spongebob},
	)

	if err != nil {
		logrus.Fatal(err)
	}

	if err := bot.Run(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("bot gracefully shutdown")
}
