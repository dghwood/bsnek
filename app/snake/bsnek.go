package snake

import (
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
		Author:     "",        // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

// Run on game Move Request
func (b *BSnek) Move(state models.GameState) models.BattlesnakeMoveResponse {
	return models.BattlesnakeMoveResponse{Move: "up"}
}
