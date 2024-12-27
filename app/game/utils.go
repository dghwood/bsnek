package game

import (
	"strings"

	"github.com/dghwood/bsnek/models"
)

var snakeLetters = [4]string{
	"A", "B", "C", "D",
}

/*
GameEngineFromString

	Generates a GameEngine from a string representation.
	eg.
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
*/
func GameEngineFromString(board string) GameEngine {
	game := GameEngine{}
	ylines := strings.Split(strings.TrimSpace(board), "\n")
	// Find snakes
	snakes := make([][]models.Coord, 4)
	numSnakes := 0
	for y := 0; y < len(ylines); y++ {
		xchars := strings.Split(strings.TrimSpace(ylines[y]), " ")
		for x := 0; x < len(xchars); x++ {
			letter := xchars[x]
			coord := models.Coord{X: x, Y: y}
			sq := game.Board.GetSquare(coord)
			if letter == "x" {
				continue
			}
			if letter == "F" {
				sq.HasFood = true
				continue
			}
			if letter == "H" {
				sq.HealthDeduction = 15
				continue
			}
			for i, sLetter := range snakeLetters {
				if strings.ToUpper(letter) == sLetter {
					if i > numSnakes {
						numSnakes = i
					}
					if strings.ToUpper(letter) == letter {
						// the head
						snakes[i] = append([]models.Coord{coord}, snakes[i]...)
					} else {
						snakes[i] = append(snakes[i], coord)
					}
					game.Board.GetSquare(coord).isBlocked = true
					break
				}
			}
		}
	}

	game.Snakes = make([]Snake, numSnakes+1)
	for i, snake := range snakes[:numSnakes+1] {
		game.Snakes[i] = Snake{
			Health: 100,
			Index:  i,
			Body:   constructSnake(snake)[0],
		}
	}
	return game
}

func constructSnake(unorderedSnake []models.Coord) [][]models.Coord {
	snake := make([][]models.Coord, 0)
	for _, dir := range Directions {
		pos := unorderedSnake[0].Add(dir)
		for i, body := range unorderedSnake {
			if i == 0 {
				// ignore the head which is positioned correctly
				continue
			}
			if pos.Equal(body) {
				newSnake := []models.Coord{body}
				for j := 1; j < len(unorderedSnake); j++ {
					if j == i {
						continue
					}
					// add unused snake body to newSnake
					newSnake = append(newSnake, unorderedSnake[j])
				}
				if len(newSnake) == 1 {
					// If there is only 1 piece left, then unorderedSnake
					// is length 2 and ordered.
					return [][]models.Coord{unorderedSnake}
				}
				for _, option := range constructSnake(newSnake) {
					snakeOption := make([]models.Coord, len(option)+1)
					snakeOption[0] = unorderedSnake[0]
					copy(snakeOption[1:], option)
					snake = append(snake, snakeOption)
				}
			}
		}
	}
	return snake
}

func GameBoardToString(game GameEngine) string {
	board := [11][11]string{}
	for y := 0; y < 11; y++ {
		for x := 0; x < 11; x++ {
			coord := models.Coord{X: x, Y: y}
			sq := game.Board.GetSquare(coord)
			if sq.HasFood {
				board[y][x] = "F"
			} else if sq.HealthDeduction > 1 {
				board[y][x] = "H"
			} else {
				board[y][x] = " "
			}
		}
	}
	// add snakes
	for i, snake := range game.Snakes {
		for _, body := range snake.Body {
			board[body.Y][body.X] = strings.ToLower(snakeLetters[i])
		}
		head := snake.Body[0]
		board[head.Y][head.X] = snakeLetters[i]
	}

	boardString := ""
	for y := 0; y < 11; y++ {
		for x := 0; x < 11; x++ {
			boardString += board[y][x]
			boardString += " "
		}
		if y == 10 {
			continue
		}
		boardString += "\n"
	}
	return boardString
}
