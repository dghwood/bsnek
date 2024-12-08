package snake

import "github.com/dghwood/bsnek/models"

type Square struct {
	HasFood         bool
	HasSnake        bool
	HasSnakeFor     int
	HealthDeduction int
}

type Board struct {
	// X, Y indexed
	board [][]Square
}

func (b *Board) Init(state models.GameState) {
	b.board = make([][]Square, state.Board.Width)
	for i := 0; i < state.Board.Width; i++ {
		b.board[i] = make([]Square, state.Board.Height)
	}

	// Add Snakes
	for _, snake := range state.Board.Snakes {
		for i, coord := range snake.Body {
			b.board[coord.X][coord.Y].HasSnake = true
			b.board[coord.X][coord.Y].HasSnakeFor = len(snake.Body) - i - 1
		}
	}

	// Add Food
	for _, coord := range state.Board.Food {
		b.board[coord.X][coord.Y].HasFood = true
	}

	// Add Hazard
	for _, coord := range state.Board.Hazards {
		b.board[coord.X][coord.Y].HealthDeduction = state.Game.Ruleset.Settings.HazardDamagePerTurn
	}
}

func (b *Board) GetSquare(x int, y int) Square {
	return b.board[x][y]
}
