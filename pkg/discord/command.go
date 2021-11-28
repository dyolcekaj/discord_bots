package discord

type SlashCommand interface {
	Command
}

type Command interface {
	Name() string
}

var _ SlashCommand = (*sCommand)(nil)
var _ Command = (*sCommand)(nil)

type sCommand struct {
	name string
}

func NewSlashCommand(name string) Command {
	return &sCommand{
		name: name,
	}
}

func (c *sCommand) Name() string {
	return c.name
}
