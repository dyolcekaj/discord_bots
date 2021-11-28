package discord

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func NewHandler(commands []Command) *discordHandler {
	cmdMap := make(map[string]Command)

	for _, cmd := range commands {
		cmdMap[cmd.Name()] = cmd
	}

	return &discordHandler{
		commands: cmdMap,
	}
}

type discordHandler struct {
	commands map[string]Command
}

var _ http.Handler = (*discordHandler)(nil)

func (h *discordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	logrus.Infof("Payload: %v", input)

	switch input.Type {
	case InputTypePing:
		ping(w)
		return
	case InputTypeCommand:
		command(w, input)
		return
	default:
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
}

func ping(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"type\": 1}")
}

func command(w http.ResponseWriter, input Input) {
	logrus.Info("handling command")
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
}
