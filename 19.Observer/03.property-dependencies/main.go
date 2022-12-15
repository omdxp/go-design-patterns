package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

type Observer interface {
	Notify(data any)
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *Person) Fire(data any) {
	for z := p.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type PropertyChange struct {
	Name  string // name of the property
	Value any    // new value
}

func (p *Person) SetAge(age int) {
	if p.age == age {
		return
	}

	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChange{"Age", age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct {
}

func (e *ElectoralRoll) Notify(data any) {
	if pc, ok := data.(PropertyChange); ok && pc.Name == "CanVote" {
		if pc.Value.(bool) {
			fmt.Println("You can vote!")
		} else {
			fmt.Println("You can't vote yet!")
		}
	}
}

func main() {
	p := NewPerson(0)
	er := &ElectoralRoll{}
	p.Subscribe(er)

	for i := 18; i < 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}

// Property dependencies are a way to express that a property depends on other properties.
