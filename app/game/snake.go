package game

import "github.com/dghwood/bsnek/models"

type Snake struct {
	Body   []models.Coord
	Health int
	Died   bool
	Index  int
}

func (s *Snake) Copy() Snake {
	return Snake{
		Body:   append([]models.Coord{}, s.Body...),
		Health: s.Health,
		Died:   s.Died,
		Index:  s.Index,
	}
}

func (s *Snake) GetHead() models.Coord {
	return s.Body[0]
}

func (s *Snake) Move(move models.Coord) {
	s.Body = append([]models.Coord{move}, s.Body[:len(s.Body)-1]...)
}

func (s *Snake) Eat() {
	s.Health = 100
	s.Body = append(s.Body, s.Body[len(s.Body)-1])
}
