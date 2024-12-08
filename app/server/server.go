package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dghwood/bsnek/models"
	"github.com/dghwood/bsnek/snake"
)

// Middleware

const ServerID = "battlesnake/github/bsnek"

func withServerID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", ServerID)
		next(w, r)
	}
}

type SnakeServer struct {
	Snake snake.BSnek
}

func (s *SnakeServer) ListenAndServe() {
	http.HandleFunc("/", withServerID(s.Index))
	http.HandleFunc("/start", withServerID(s.Start))
	http.HandleFunc("/move", withServerID(s.Move))
	http.HandleFunc("/end", withServerID(s.End))

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	log.Printf("Running Battlesnake at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func (s *SnakeServer) Index(w http.ResponseWriter, r *http.Request) {
	response := s.Snake.Info()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode info response, %s", err)
	}
}

func (s *SnakeServer) Start(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode start json, %s", err)
		return
	}
	s.Snake.Start(state)
}

func (s *SnakeServer) Move(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode move json, %s", err)
		return
	}

	response := s.Snake.Move(state)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode move response, %s", err)
		return
	}
}

func (s *SnakeServer) End(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode end json, %s", err)
		return
	}
	s.Snake.End(state)
}
