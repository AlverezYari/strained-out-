package game

import (
	"encoding/json"
	"net/http"
)

// StateHandler returns an http.HandlerFunc that serves the game state as JSON.
func StateHandler(mgr *Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(mgr.State()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
