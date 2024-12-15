package game

/* GameEngine

This is basically a game engine for BattleSnake
*/

import (
	"github.com/dghwood/bsnek/models"
)

type Snake struct {
	Body   []models.Coord
	Health int
	Died   bool
}

func (s *Snake) GetHead() models.Coord {
	return s.Body[0]
}
func (s *Snake) GetDirections() [4]models.Coord {
	moves := [4]models.Coord{}
	head := s.GetHead()
	for i, direction := range Directions {
		moves[i] = head.Add(direction)
	}
	return moves
}

func (s *Snake) Move(move models.Coord) {
	s.Body = append([]models.Coord{move}, s.Body[:len(s.Body)-1]...)
}

func (s *Snake) Eat() {
	s.Body = append(s.Body, s.Body[len(s.Body)-1])
}

type Square struct {
	HasFood         bool
	isBlocked       bool
	HealthDeduction int
}

type GameBoard struct {
	board [11][11]Square
}

func (b *GameBoard) GetSquare(coord models.Coord) *Square {
	if coord.X < 0 || coord.Y < 0 || coord.X > 10 || coord.Y > 10 {
		return &Square{isBlocked: true}
	}
	return &b.board[coord.X][coord.Y]
}

type GameEngine struct {
	Board  GameBoard
	Snakes []Snake
}

func (g *GameEngine) HandleHead2HeadCollision(snake1, snake2 *Snake) {
	// Note: this might get called multiple times for numerous H2H collisions
	// in the same spot.
	// snake1 := &g.Snakes[i]
	// snake2 := &g.Snakes[j]
	cond := len(snake1.Body) - len(snake2.Body)
	if cond > 0 {
		// snake1 bigger
		snake2.Died = true
	} else if cond < 0 {
		// snake2 bidder
		snake1.Died = true
	} else if cond == 0 {
		// both die
		snake1.Died = true
		snake2.Died = true
	}
}

func (g *GameEngine) Init(state models.GameState) {
	// b.InitEmpty(state.Board.Width, state.Board.Height)

	// Add Snakes
	g.Snakes = make([]Snake, len(state.Board.Snakes))
	for i, snake := range state.Board.Snakes {
		g.Snakes[i] = Snake{
			Body:   snake.Body,
			Health: snake.Health,
		}
		for _, coord := range snake.Body[:len(snake.Body)-1] {
			// Don't add the tail
			g.Board.GetSquare(coord).isBlocked = true
		}
	}

	// Add Food
	for _, coord := range state.Board.Food {
		g.Board.GetSquare(coord).HasFood = true
	}

	// Add Hazard
	for _, coord := range state.Board.Hazards {
		g.Board.GetSquare(coord).HealthDeduction = state.Game.Ruleset.Settings.HazardDamagePerTurn
	}
}

func (g *GameEngine) PlayScenario(moves []models.Coord) {
	// FYI: Moves are indexed by snake

	for i := 0; i < len(moves); i++ {
		move1 := moves[i]
		snake := &g.Snakes[i]

		// Check no moves collide
		for j := i + 1; j < len(moves); j++ {
			move2 := moves[j]
			if move1.Equal(move2) {
				// If the snake is dead you can probably skip
				g.HandleHead2HeadCollision(snake, &g.Snakes[j])
			}
		}

		sq := g.Board.GetSquare(move1)

		// Check other collisons
		if sq.isBlocked {
			snake.Died = true
		}

		// Assign Food
		if sq.HasFood {
			sq.HasFood = false
			snake.Eat()
		}

		// Check Health
		snake.Health -= 1 + sq.HealthDeduction
		if snake.Health <= 0 {
			snake.Died = true
		}

		snake.Move(move1)
	}

	// You have to wait for all the turns before handling
	// dead snakes.

}

var Directions = [4]models.Coord{
	{X: 0, Y: 1},  // up
	{X: 0, Y: -1}, // down
	{X: -1, Y: 0}, // left
	{X: 1, Y: 0},  // right
}

// func (g *GameEngine) IdentifyMoves() [][]models.Coord {
// 	moves := make([][]models.Coord, len(g.Snakes)*4)
// 	for i := 0; i < len(g.Snakes); i++ {
// 		snake := g.Snakes[i]
// 		directions := snake.GetDirections()
// 	}

// 	for i := 0; i < len(g.Snakes) * 4; i++ {

// 	}
// }
