package game

/* GameEngine

This is basically a game engine for BattleSnake
*/

import (
	"math/rand"

	"github.com/dghwood/bsnek/models"
	"github.com/dghwood/bsnek/permutations"
)

var Directions = [4]models.Coord{
	{X: 0, Y: 1},  // up
	{X: 0, Y: -1}, // down
	{X: -1, Y: 0}, // left
	{X: 1, Y: 0},  // right
}

type GameEngine struct {
	Board    GameBoard
	Snakes   []Snake
	Moves    []models.Coord
	MoveNum  int
	YouIndex int
}

func (g *GameEngine) copySnakes() []Snake {
	snakes := make([]Snake, len(g.Snakes))
	for i := 0; i < len(g.Snakes); i++ {
		snakes[i] = g.Snakes[i].Copy()
	}
	return snakes
}

func (g *GameEngine) Copy() GameEngine {
	return GameEngine{
		Board:   g.Board.Copy(),
		Snakes:  g.copySnakes(),
		Moves:   append([]models.Coord{}, g.Moves...),
		MoveNum: g.MoveNum,
	}
}

func (g *GameEngine) Init(state models.GameState) {
	// Add Turn
	turn := state.Turn
	g.Board.Turn = turn

	// Add Snakes
	g.Snakes = make([]Snake, len(state.Board.Snakes))
	for i, snake := range state.Board.Snakes {
		g.Snakes[i] = Snake{
			Body:   snake.Body, // Copy?
			Health: snake.Health,
			Index:  i,
		}
		for j, coord := range snake.Body[:len(snake.Body)-1] {
			// Don't add the tail, or do this in reverse
			// Since snake tails can overlap
			g.Board.SetBlockedUntil(coord, len(snake.Body), j)
		}
		// Find your index
		if state.You.ID == snake.ID {
			g.YouIndex = i
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

func (g *GameEngine) handleHead2HeadCollision(snake1, snake2 *Snake) {
	// Note: this might get called multiple times for numerous H2H collisions
	// in the same spot.
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

func (g *GameEngine) PlayScenario(moves []models.Coord) {
	deadIndex := make([]int, 0)
	aliveIndex := make([]int, 0)

	for i := 0; i < len(moves); i++ {
		move1 := moves[i]
		snake := &g.Snakes[i]

		// Check no moves collide
		for j := i + 1; j < len(moves); j++ {
			move2 := moves[j]
			if move1.Equal(move2) {
				// If the snake is dead you can probably skip
				g.handleHead2HeadCollision(snake, &g.Snakes[j])
			}
		}

		// Check other collisons
		if g.Board.IsBlocked(move1) {
			snake.Died = true
		}

		sq := g.Board.GetSquare(move1)
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

		if snake.Died {
			deadIndex = append(deadIndex, i)
		} else {
			aliveIndex = append(aliveIndex, i)
		}
	}

	// Now move snakes
	for _, i := range deadIndex {
		// Dead first, so for h2h collisions you overwrite the blockedUtilTurn
		snake := g.Snakes[i]
		for _, body := range snake.Body {
			g.Board.GetSquare(body).BlockedUtilTurn = g.Board.Turn
		}
	}
	for _, i := range aliveIndex {
		move := moves[i]
		g.Snakes[i].Move(move)
		g.Board.SetBlockedUntil(move, len(g.Snakes[i].Body), 0)
	}
}

func (g *GameEngine) PlayRandomScenario() {
	scenarios := g.GetAllScenarios()
	index := rand.Intn(len(scenarios))
	g.PlayScenario(scenarios[index])
}

func (g *GameEngine) GetScenarios() ([][]models.Coord, []int) {
	scenarios := make([][]models.Coord, len(g.Snakes))
	lens := make([]int, len(g.Snakes))
	for i := 0; i < len(g.Snakes); i++ {
		head := g.Snakes[i].GetHead()
		moves := make([]models.Coord, 0)
		for _, direction := range Directions {
			move := head.Add(direction)
			// TODO(duncanwood): Check board is updated.. since blocked might be stale
			if !g.Board.IsBlocked(move) {
				moves = append(moves, head.Add(direction))
			}
		}
		if len(moves) > 0 {
			scenarios[i] = moves
			lens[i] = len(moves)
		} else {
			// When no valid moves just pick up
			scenarios[i] = []models.Coord{Directions[0]}
			lens[i] = 1
		}

	}
	return scenarios, lens
}

func (g *GameEngine) GetAllScenarios() [][]models.Coord {
	scenarios, lens := g.GetScenarios()
	// Note: update Permutations to assume >=1 lens
	perms := permutations.Permutations(lens)
	moves := make([][]models.Coord, len(perms))
	for i, indexes := range perms {
		moves[i] = make([]models.Coord, len(indexes))
		for j := 0; j < len(indexes); j++ {
			moves[i][j] = scenarios[j][indexes[j]]
		}
	}
	return moves
}

/* Scoring */

type ScoredSquare struct {
	BlockedUtilTurn int
}
type ScoredBoard struct {
	board [11][11]ScoredSquare
}

func (s *ScoredBoard) GetSquare(coord models.Coord) *ScoredSquare {
	return &s.board[coord.X][coord.Y]
}

func (g *GameEngine) Score() float32 {
	if g.Snakes[g.YouIndex].Died {
		return -1.0
	}
	scoredBoard := ScoredBoard{}
	// add snakes to board

	// Keep score of squares
	snakeSquares := make([]int, len(g.Snakes))

	for turn := 0; turn < 121; turn++ {
		for i, snake := range g.Snakes {
			head := snake.GetHead()
			for _, dir := range Directions {
				coord := head.Add(dir)
				sq := scoredBoard.GetSquare(coord)
				if turn <= sq.BlockedUtilTurn {
					// square blocked
					continue
				}
				sq.BlockedUtilTurn = turn + len(snake.Body)
				snakeSquares[i] += 1
			}
		}
	}
	return 0.0
}
