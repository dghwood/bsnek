package snake

import (
	"testing"

	"github.com/dghwood/bsnek/models"
)

func TestBoardFood(t *testing.T) {
	state := models.GameState{
		Board: models.Board{
			Width:  11,
			Height: 11,
			Food: []models.Coord{
				{X: 1, Y: 1},
				{X: 9, Y: 7},
			},
		},
	}
	board := Board{}
	board.Init(state)
	if !board.GetSquareFromXY(1, 1).HasFood {
		t.Fatalf("Square 1,1 does not have food")
	}
	if !board.GetSquareFromXY(9, 7).HasFood {
		t.Fatalf("Square 9, 7 does not have food")
	}
}

func TestBoardSnake(t *testing.T) {
	state := models.GameState{
		Board: models.Board{
			Width:  11,
			Height: 11,
			Snakes: []models.Battlesnake{
				{Body: []models.Coord{
					{X: 1, Y: 1},
					{X: 1, Y: 2},
					{X: 1, Y: 3},
				}},
			},
		},
	}
	board := Board{}
	board.Init(state)
	if !board.GetSquareFromXY(1, 1).HasSnake {
		t.Fatalf("Square 1,1 does not have a snake")
	}
	if !board.GetSquareFromXY(1, 2).HasSnake {
		t.Fatalf("Square 9, 7 does not have a snake")
	}
	if board.GetSquareFromXY(1, 1).HasSnakeFor != 2 {
		t.Fatalf("Square 1, 1 HasSnakeFor != 2, equals: %d", board.GetSquareFromXY(1, 1).HasSnakeFor)
	}
}

func TestBoardHazard(t *testing.T) {
	state := models.GameState{
		Game: models.Game{
			Ruleset: models.Ruleset{
				Settings: models.RulesetSettings{
					HazardDamagePerTurn: 3,
				},
			},
		},
		Board: models.Board{
			Width:  11,
			Height: 11,
			Hazards: []models.Coord{
				{X: 1, Y: 1},
				{X: 2, Y: 2},
			},
		},
	}
	board := Board{}
	board.Init(state)
	if board.GetSquareFromXY(1, 1).HealthDeduction != 3 {
		t.Fatalf("Square 1,1 health deduction is not 3: equals: %d",
			board.GetSquareFromXY(1, 1).HealthDeduction)
	}
	if board.GetSquareFromXY(10, 10).HealthDeduction != 0 {
		t.Fatalf("Square 10, 10 health deduction is not 0: equals %d",
			board.GetSquareFromXY(10, 10).HealthDeduction)
	}
}
