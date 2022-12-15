package main

import "fmt"

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

func PlayGame(game Game) {
	game.Start()
	for !game.HaveWinner() {
		game.TakeTurn()
	}
	fmt.Printf("Player %d wins.\n", game.WinningPlayer())
}

type Chess struct {
	turn, currentPlayer int
	maxTurns            int
}

func NewGameOfChess() *Chess {
	return &Chess{1, 0, 10}
}

func (c *Chess) Start() {
	fmt.Println("Starting a game of chess")
}

func (c *Chess) TakeTurn() {
	fmt.Println("Turn", c.turn, "taken by player", c.currentPlayer)
	c.turn++
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *Chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *Chess) WinningPlayer() int {
	return c.currentPlayer
}

func main() {
	chess := NewGameOfChess()
	PlayGame(chess)
}

// Template Method is a behavioral design pattern that defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure.
