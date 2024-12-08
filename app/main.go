package main

import (
	"github.com/dghwood/bsnek/server"
	"github.com/dghwood/bsnek/snake"
)

func main() {
	server := server.SnakeServer{Snake: snake.BSnek{}}
	server.ListenAndServe()
}
