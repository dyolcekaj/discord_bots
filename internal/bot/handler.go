package bot

import "net/http"

type discordHandler struct {
}

var _ http.Handler = (*discordHandler)(nil)

func (h *discordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func ping(w http.ResponseWriter, r *http.Request) {

}
