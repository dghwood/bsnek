package game

import (
	"fmt"
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

func TestGetAllScenarios(t *testing.T) {
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
					{X: 0, Y: 5},
					{X: 1, Y: 5},
					{X: 2, Y: 5},
				},
			},
		},
	}
	// scenarios, lens := game.GetScenarios()
	moves := game.GetAllScenarios()
	e := [][]models.Coord{
		{{X: 0, Y: 1}, {X: 0, Y: 6}},
		{{X: 0, Y: 1}, {X: 0, Y: 4}},
		{{X: 0, Y: 1}, {X: 1, Y: 5}},

		{{X: 1, Y: 0}, {X: 0, Y: 6}},
		{{X: 1, Y: 0}, {X: 0, Y: 4}},
		{{X: 1, Y: 0}, {X: 1, Y: 5}},
	}

	if len(moves) != len(e) {
		t.Fatalf("Moves are not the expected length: actual %d, expected: %d", len(moves), len(e))
	}
	for _, emove := range e {
		foundMove := false
		for _, move := range moves {
			if move[0].Equal(emove[0]) && move[1].Equal(emove[1]) {
				foundMove = true
			}
		}
		if !foundMove {
			t.Fatal("Not able to find expected move", emove)
		}
	}
}

func TestGetAllScenariosWithNoValidMoves(t *testing.T) {
	// TODO(duncanwood): Make a test to check what happens
	// when a snake doesn't have valid moves
	// Should move up, and nothing break
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
					{X: 0, Y: 5},
					{X: 1, Y: 5},
					{X: 2, Y: 5},
				},
			},
		},
	}
	// scenarios, lens := game.GetScenarios()
	moves := game.GetAllScenarios()
	e := [][]models.Coord{
		{{X: 0, Y: 1}, {X: 0, Y: 6}},
		{{X: 0, Y: 1}, {X: 0, Y: 4}},
		{{X: 0, Y: 1}, {X: 1, Y: 5}},

		{{X: 1, Y: 0}, {X: 0, Y: 6}},
		{{X: 1, Y: 0}, {X: 0, Y: 4}},
		{{X: 1, Y: 0}, {X: 1, Y: 5}},
	}

	if len(moves) != len(e) {
		t.Fatalf("Moves are not the expected length: actual %d, expected: %d", len(moves), len(e))
	}
	for _, emove := range e {
		foundMove := false
		for _, move := range moves {
			if move[0].Equal(emove[0]) && move[1].Equal(emove[1]) {
				foundMove = true
			}
		}
		if !foundMove {
			t.Fatal("Not able to find expected move", emove)
		}
	}
}

func TestCopy(t *testing.T) {
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

	coord := models.Coord{X: 0, Y: 0}

	game.Board.GetSquare(coord).HasFood = true

	newGame := game.Copy()

	game.Snakes[0].Body[0].X = 1
	game.Board.GetSquare(coord).HasFood = false

	if newGame.Snakes[0].Body[0].X != 0 {
		t.Fatalf("Snake coord not copied: %d", newGame.Snakes[0].Body[0].X)
	}

	if newGame.Board.GetSquare(coord).HasFood != true {
		t.Fatalf("Board not copied")
	}

}

func TestUtilExample(t *testing.T) {
	board := `
	x x x x x x x x x x x
	x x x x x x x x x x x
	x A x x x x x x x x x
	x a x x x x x x x x x
	x a x x x x x x x x x
	a a x x x F x x x x x
	a a a x x x x B x x x
	a a a x x x x b x x x
	x x x x x x x b x x x
	x x x x x x x b b x x
	x x x x x x x x x x x
	x x x x x x x x x x x
	`
	game := GameEngineFromString(board)
	for i := 0; i < 10; i++ {
		game.PlayRandomScenario()
		fmt.Println(GameBoardToString(game))
		fmt.Println("--")
	}

}
