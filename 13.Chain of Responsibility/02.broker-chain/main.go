package main

import (
	"fmt"
	"sync"
)

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(*Query)
}

type Observable interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Fire(query *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(observer Observer) {
	g.observers.Store(observer, struct{}{})
}

func (g *Game) Unsubscribe(observer Observer) {
	g.observers.Delete(observer)
}

func (g *Game) Fire(query *Query) {
	g.observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(query)
		return true
	})
}

type Creature struct {
	game            *Game // Mediator
	Name            string
	attack, defense int
}

func NewCreature(game *Game, name string, attack, defense int) *Creature {
	return &Creature{game: game, Name: name, attack: attack, defense: defense}
}

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(), c.Defense())
}

type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func (c *CreatureModifier) Handle(q *Query) {
	// nothing
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(game *Game, creature *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{CreatureModifier{game, creature}}
	game.Subscribe(d)
	return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == d.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

func main() {
	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin.String())

	{
		dam := NewDoubleAttackModifier(game, goblin)
		defer dam.Close()
		fmt.Println(goblin.String())
	}

	fmt.Println(goblin.String())
}

// Broker Chain of Responsibility Pattern
// CoR, Mediator, Observer, CQS
