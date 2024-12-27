package game

import (
	"fmt"
	"time"
)

/*
	Runner

# So what I want is to have a situation like this

Y1 -> (O1, O2)
Y2 -> (O1, O2)

--> Then you score the moves & potentially prune

Y11 -> (O11, O12, O21, O22)
Y12 -> (O11, O12, O21, O22)

-- Then you score
*/

func Runner(queue chan *GameEngine, quit chan bool) {
	num_games := 0
	max_moves := 0
runner:
	for {
		select {
		case <-quit:
			fmt.Println("Quitting", num_games, max_moves)
			break runner
		case game := <-queue:
			if game.MoveNum > max_moves {
				max_moves = game.MoveNum
			}
			moves := game.GetAllScenarios()
			games := make([]GameEngine, len(moves))
			for i, move := range moves {
				g := game.Copy()
				g.PlayScenario(move)
				games[i] = g
				num_games += 1
				// Score
				// Send new ones to runner
			}
			// Send the new games
			for _, g := range games {
				if g.Score() > -1 {
					// This can get blocked
					queue <- &g
				}
			}
		default:
			// Oh this fires when there isn't game in the queue as well as on blocked
			fmt.Println("Blocked channel", num_games)
		}
	}
}

func StartRun(game *GameEngine) {
	queue := make(chan *GameEngine, 1000000)
	quit := make(chan bool)
	num_cores := 1

	for i := 0; i < num_cores; i++ {
		go Runner(queue, quit)
	}
	queue <- game
	time.Sleep(time.Duration(100 * time.Millisecond))
	fmt.Println("sending quit")
	for i := 0; i < num_cores; i++ {
		quit <- true
	}
}
