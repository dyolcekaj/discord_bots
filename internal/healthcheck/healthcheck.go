package healthcheck

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Up())
}

func Up() Response {
	return Response{
		Status: "UP",
		Reason: "",
	}
}
