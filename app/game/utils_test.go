package game

import (
	"fmt"
	"testing"
)


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
	`
	state := GameStateFromString(board)
	game := GameEngine{}
	game.Init(state)
	fmt.Println(GameBoardToString(game))
}
