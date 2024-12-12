package http

import (
	"encoding/json"
	"net/http"
)

func (s *Server) HandleStart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	startElements := []Response{
		{
			Result: "Fire",
			Emoji:  "🔥",
			IsNew:  false,
		},
		{
			Result: "Water",
			Emoji:  "💧",
			IsNew:  false,
		},
		{
			Result: "Earth",
			Emoji:  "🌍",
			IsNew:  false,
		},
		{
			Result: "Air",
			Emoji:  "💨",
			IsNew:  false,
		},
	}

	if err := json.NewEncoder(w).Encode(startElements); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
