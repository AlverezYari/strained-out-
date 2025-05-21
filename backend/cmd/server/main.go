package main

import (
	"log"
	"net/http"

	"github.com/example/imperial/internal/game"
)

func main() {
	mgr := game.NewManager()

	mux := http.NewServeMux()
	mux.Handle("/api/state", game.StateHandler(mgr))

	log.Println("server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
