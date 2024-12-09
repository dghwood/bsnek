package snake

import "github.com/dghwood/bsnek/models"

type Square struct {
	HasFood         bool
	HasSnake        bool
	HasSnakeFor     int
	HealthDeduction int
	isError         bool
}

func (sq Square) isBlocked() bool {
	return (sq.HasSnake || sq.isError)
}

type Board struct {
	// X, Y indexed
	board  [][]Square
	Height int
	Width  int
}

func (b *Board) InitEmpty(width int, height int) {
	b.board = make([][]Square, width)
	for i := 0; i < width; i++ {
		b.board[i] = make([]Square, height)
	}
	b.Height = height
	b.Width = width
}

func (b *Board) Init(state models.GameState) {
	b.InitEmpty(state.Board.Width, state.Board.Height)

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

func (b *Board) GetSquareFromXY(x int, y int) Square {
	return b.board[x][y]
}

func (b *Board) GetSquare(coord models.Coord) Square {
	if coord.X < 0 || coord.Y < 0 || coord.X >= b.Width || coord.Y >= b.Height {
		return Square{isError: true}
	}
	return b.board[coord.X][coord.Y]
}
