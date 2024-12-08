package snake

import (
	"github.com/dghwood/bsnek/models"
)

type BSnek struct {
}

func (b *BSnek) Start(state models.GameState) {

}

func (b *BSnek) End(state models.GameState) {

}

func (b *BSnek) Info() models.BattlesnakeInfoResponse {
	return models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "",        // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

func (b *BSnek) Move(state models.GameState) models.BattlesnakeMoveResponse {
	return models.BattlesnakeMoveResponse{Move: "up"}
}
