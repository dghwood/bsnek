package snake

import (
	"fmt"
	"sort"

	"github.com/dghwood/bsnek/models"
)

type BSnek struct {
}

// Run on game Start
func (b *BSnek) Start(state models.GameState) {

}

// Run on game End
func (b *BSnek) End(state models.GameState) {

}

// Run to get Snake information
func (b *BSnek) Info() models.BattlesnakeInfoResponse {
	return models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "dghwood", // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

// Run on game Move Request
func (b *BSnek) Move(state models.GameState) models.BattlesnakeMoveResponse {
	fmt.Print("Move")
	board := Board{}
	board.Init(state)
	head := state.You.Head
	scoredMoves := b.EvaulateMoves(head, board)
	finalMove := scoredMoves[0]
	return models.BattlesnakeMoveResponse{Move: BackoutDirection(head, finalMove.Coord)}
}

/* Move Evaulation */

func (b *BSnek) EvaulateMove(head models.Coord, board Board) ScoredMove {
	scoredMove := ScoredMove{Coord: head}

	if board.GetSquare(head).isBlocked() {
		scoredMove.Score = -1
		return scoredMove
	}

	return scoredMove
}

func (b *BSnek) EvaulateMoves(head models.Coord, board Board) []ScoredMove {
	scoredMoves := make([]ScoredMove, 0)
	for _, direction := range Directions {
		newCoord := head.Add(direction)
		scoredMoves = append(scoredMoves, b.EvaulateMove(newCoord, board))
	}
	sort.Slice(scoredMoves, func(i, j int) bool {
		return scoredMoves[j].Score < scoredMoves[i].Score
	})
	return scoredMoves
}
