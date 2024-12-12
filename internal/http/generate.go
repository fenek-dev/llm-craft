package http

import (
	"encoding/json"
	"net/http"
)

func (s *Server) HandleGenerate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	el1 := r.URL.Query().Get("first")
	el2 := r.URL.Query().Get("second")

	result, isNew, err := s.s.Generate(r.Context(), el1, el2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := Response{
		Result: result.Name,
		Emoji:  result.Emoji,
		IsNew:  isNew,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
