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

func TestValidMoves(t *testing.T) {
	board := Board{}
	board.InitEmpty(11, 11)
	head := models.Coord{X: 0, Y: 0}
	snake := BSnek{}
	moves := snake.EvaulateMoves(head, board)
	if len(moves) != 4 {
		t.Fatalf("Moves is not length 4, length %d", len(moves))
	}
	if moves[0].Coord.X < 0 || moves[1].Coord.Y < 0 {
		t.Fatal("Top move is invalid", moves[0].Coord)
	}
	direction := BackoutDirection(head, moves[0].Coord)
	if direction == "up" || direction == "left" {
		t.Fatalf("Wrong direction: %s", direction)
	}
}
