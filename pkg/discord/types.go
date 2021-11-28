package discord

type InputType int

const (
	InputTypePing InputType = iota + 1
	InputTypeCommand
)

type Member struct {
	User     User     `json:"user"`
	Roles    []string `json:"roles"`
	Nickname string   `json:"nick"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Input struct {
	Type      InputType   `json:"type"`
	Token     string      `json:"token"`
	Member    Member      `json:"member"`
	ID        string      `json:"id"`
	GuildID   string      `json:"guild_id"`
	ChannelID string      `json:"channel_id"`
	Data      interface{} `json:"data"`
}
