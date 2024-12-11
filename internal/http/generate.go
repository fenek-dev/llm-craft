package http

import (
	"encoding/json"
	"net/http"
)

func (s *Server) HandleGenerate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	el1 := r.URL.Query().Get("first")
	el2 := r.URL.Query().Get("second")

	result, err := s.s.Generate(r.Context(), el1, el2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
