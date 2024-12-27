package game

import "github.com/dghwood/bsnek/models"

const boardHeight = 11
const boardWidth = 11

type Square struct {
	HasFood         bool
	isBlocked       bool
	HealthDeduction int
}

type GameBoard struct {
	board [boardWidth][boardHeight]Square
}

func (b *GameBoard) GetSquare(coord models.Coord) *Square {
	if coord.X < 0 || coord.Y < 0 || coord.X > boardWidth || coord.Y > boardHeight {
		return &Square{isBlocked: true}
	}
	return &b.board[coord.X][coord.Y]
}

func (b *GameBoard) Copy() GameBoard {
	// If this just copies the board, then not sure I need it.
	return GameBoard{
		board: b.board,
	}
}
