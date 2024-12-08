package snake

import (
	"testing"

	"github.com/dghwood/bsnek/models"
)

func TestValidMove(t *testing.T) {
	validMoves := map[string]bool{"up": true, "down": true, "left": true, "right": true}
	state := models.GameState{}
	snake := BSnek{}
	response := snake.Move(state)
	_, ok := validMoves[response.Move]
	if !ok {
		t.Fatalf("Response not valid move: %s", response.Move)
	}
}
