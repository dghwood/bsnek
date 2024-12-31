package game

import (
	"testing"

	"github.com/dghwood/bsnek/models"
)

func TestSnakeMove(t *testing.T) {
	snake := Snake{
		Body: []models.Coord{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
		},
	}

	move := models.Coord{X: 0, Y: 1}
	snake.Move(move)
	if !snake.Body[0].Equal(move) {
		t.Fatal("Snake Body didn't move", snake.Body)
	}
	if !snake.Body[1].Equal(models.Coord{X: 0, Y: 0}) {
		t.Fatal("Snake Body 1 didn't move", snake.Body)
	}
	if !snake.Body[2].Equal(models.Coord{X: 1, Y: 0}) {
		t.Fatal("Snake Body 2 didn't move", snake.Body)
	}
	if len(snake.Body) != 3 {
		t.Fatal("Snake is too long", snake.Body)
	}
}
