package bot

import (
	"github.com/bwmarrin/discordgo"
)

type Bot interface {
}

type bot struct {
	s *discordgo.Session
}
