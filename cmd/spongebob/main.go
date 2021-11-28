package main

import "github.com/dyolcekaj/discord_bots/internal/bot"

func init() {
	// do stuff
}

func main() {
	bot := bot.New(bot.BotOptions{})

	<-bot.Shutdown()
}
