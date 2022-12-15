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
	p.age = age
	p.Fire(PropertyChange{"Age", age})
}

func (p *Person) Age() int {
	return p.age
}

type TrafficManagement struct {
	o *Observable
}

func NewTrafficManagement(o *Observable) *TrafficManagement {
	return &TrafficManagement{o: o}
}

func (t *TrafficManagement) Notify(data any) {
	if pc, ok := data.(PropertyChange); ok && pc.Name == "Age" {
		if pc.Value.(int) < 18 {
			fmt.Println("You are too young to drive")
		} else {
			fmt.Println("You are old enough to drive")
			t.o.Unsubscribe(t)
		}
	}
}

func main() {
	p := NewPerson(15)
	tm := NewTrafficManagement(&p.Observable)
	p.Subscribe(tm)
	for i := 16; i < 25; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}

// Property observers are a way to observe changes to a property.
