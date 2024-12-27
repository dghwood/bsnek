package game

import (
	"testing"

	"github.com/dghwood/bsnek/models"
)

func TestRunner(t *testing.T) {
	game := GameEngine{
		Snakes: []Snake{
			{
				Health: 100,
				Body: []models.Coord{
					{X: 0, Y: 0},
					{X: 1, Y: 0},
					{X: 2, Y: 0},
				},
				Index: 0,
			},
			{
				Health: 100,
				Body: []models.Coord{
					{X: 0, Y: 9},
					{X: 1, Y: 9},
					{X: 2, Y: 9},
				},
				Index: 1,
			},
		},
	}

	StartRun(&game)
}
