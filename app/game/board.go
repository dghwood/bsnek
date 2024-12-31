package game

import "github.com/dghwood/bsnek/models"

const boardHeight = 11
const boardWidth = 11

type Square struct {
	HasFood         bool
	BlockedUtilTurn int
	HealthDeduction int
}

type GameBoard struct {
	board [boardWidth][boardHeight]Square
	Turn  int
}

func (b *GameBoard) IsBlocked(coord models.Coord) bool {
	// Making this > since Turn 0 is a the default, and is causing
	// lots of issues
	return b.GetSquare(coord).BlockedUtilTurn > b.Turn
}

func (b *GameBoard) GetSquare(coord models.Coord) *Square {
	if coord.X < 0 || coord.Y < 0 || coord.X >= boardWidth || coord.Y >= boardHeight {
		return &Square{BlockedUtilTurn: 99999}
	}
	return &b.board[coord.X][coord.Y]
}

func (b *GameBoard) Copy() GameBoard {
	// If this just copies the board, then not sure I need it.
	return GameBoard{
		board: b.board,
	}
}

func (b *GameBoard) SetBlockedUntil(coord models.Coord, snakeLength int, bodyIndex int) {
	// The last body segment should be BlockedUntil the current turn
	// snakeLength - 1 == bodyIndex
	b.GetSquare(coord).BlockedUtilTurn = b.Turn + snakeLength - 1 - bodyIndex
}
