package discord

import "github.com/sirupsen/logrus"

const commandURLFormat = "https://discord.com/api/v8/applications/%s/commands"

func (b *bot) registerOrUpdateCommands() error {
	for _, cmd := range b.commands {
		logrus.Infof("Registering or updating %s", cmd.Name())

		if err := b.registerCommand(cmd); err != nil {
			return nil
		}
	}

	return nil
}

func (b *bot) registerCommand(cmd Command) error {
	return nil
}
