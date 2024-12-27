package game

import (
	"fmt"
	"testing"
)

func TestGameEngineFromString(t *testing.T) {
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
	GameEngineFromString(board)
}

func TestGameEngineToString(t *testing.T) {
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
	fmt.Println(GameBoardToString(game))
}
