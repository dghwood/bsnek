package game

import (
	"testing"

	"github.com/dghwood/bsnek/models"
)

func TestMove(t *testing.T) {
	game := GameEngine{
		Snakes: []Snake{
			{
				Body: []models.Coord{
					{X: 0, Y: 0},
					{X: 1, Y: 0},
					{X: 2, Y: 0},
				},
			},
			{
				Body: []models.Coord{
					{X: 10, Y: 0},
					{X: 9, Y: 0},
					{X: 8, Y: 0},
				},
			},
		},
	}

	game.PlayScenario([]models.Coord{{X: 0, Y: 1}, {X: 10, Y: 1}})
	expectedBody1 := []models.Coord{
		{X: 0, Y: 1},
		{X: 0, Y: 0},
		{X: 1, Y: 0},
	}
	expectedBody2 := []models.Coord{
		{X: 10, Y: 1},
		{X: 10, Y: 0},
		{X: 9, Y: 0},
	}
	for i, body := range game.Snakes[0].Body {
		if !body.Equal(expectedBody1[i]) {
			t.Fatal("Snake Body moved incorrectly", body, expectedBody1[i])
		}
	}

	for i, body := range game.Snakes[1].Body {
		if !body.Equal(expectedBody2[i]) {
			t.Fatal("Snake Body moved incorrectly", body, expectedBody2[i])
		}
	}
}

func TestEat(t *testing.T) {
	game := GameEngine{
		Snakes: []Snake{
			{
				Body: []models.Coord{
					{X: 0, Y: 0},
					{X: 1, Y: 0},
					{X: 2, Y: 0},
				},
			},
		},
	}
	move := models.Coord{X: 0, Y: 1}
	game.Board.GetSquare(move).HasFood = true

	game.PlayScenario([]models.Coord{move})
	expectedBody := []models.Coord{
		{X: 0, Y: 1},
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
	}
	if len(game.Snakes[0].Body) != len(expectedBody) {
		t.Fatalf("Snake isn't the correct len, expected 4, got %d", len(game.Snakes[0].Body))
	}
	for i, body := range game.Snakes[0].Body {
		if !body.Equal(expectedBody[i]) {
			t.Fatal("Snake Body moved incorrectly", body, expectedBody[i])
		}
	}
}
