package snake

import "github.com/dghwood/bsnek/models"

// Valid Directions of movement
var DirectionStrings = [4]string{
	"up",
	"down",
	"left",
	"right",
}

var Directions = [4]models.Coord{
	{X: 0, Y: 1},  // up
	{X: 0, Y: -1}, // down
	{X: -1, Y: 0}, // left
	{X: 1, Y: 0},  // right
}

func BackoutDirection(head models.Coord, move models.Coord) string {
	proposedDirection := head.Minus(move)
	for i, direction := range Directions {
		if proposedDirection.X == direction.X && proposedDirection.Y == direction.Y {
			return DirectionStrings[i]
		}
	}
	return "unknown"
}
