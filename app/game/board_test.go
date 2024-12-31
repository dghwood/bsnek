package game

import (
	"testing"

	"github.com/dghwood/bsnek/models"
)

func TestOffBoard(t *testing.T) {
	board := GameBoard{}
	coord := models.Coord{X: -1, Y: 0}
	if !board.IsBlocked(coord) {
		t.Fatal("Blocked not working off board")
	}
}

func TestOnBoard(t *testing.T) {
	board := GameBoard{}
	coord := models.Coord{X: 0, Y: 0}
	if board.IsBlocked(coord) {
		t.Fatal("Board is blocked, and it shouldn't be")
	}
}

func TestBlockedUntil(t *testing.T) {
	board := GameBoard{Turn: 1}
	coord := models.Coord{X: 0, Y: 0}
	board.SetBlockedUntil(coord, 3, 2)
	if board.GetSquare(coord).BlockedUtilTurn != board.Turn {
		t.Fatal("Blocked Until set incorrectly")
	}
}
