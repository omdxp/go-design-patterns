package main

type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		println("flying")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int {
	return l.age
}

func (l *Lizard) SetAge(age int) {
	l.age = age
}

func (l *Lizard) Crawl() {
	if l.age < 10 {
		println("crawling")
	}
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) Age() int {
	return d.bird.Age()
}

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func NewDragon() *Dragon {
	return &Dragon{
		bird:   Bird{},
		lizard: Lizard{},
	}
}

func main() {
	d := NewDragon()
	d.SetAge(5)
	d.Fly()
	d.Crawl()
}

// Multiple aggregation is a design pattern that allows us to compose multiple objects into a single object.
// This is useful when we want to add functionality to an object without modifying its source code.
// In this example, we have a Bird and a Lizard that can both fly and crawl.
// We want to create a Dragon that can fly and crawl, but we don't want to modify the Bird and Lizard structs.
// We can do this by creating a Dragon struct that has a Bird and a Lizard as fields.
// The Dragon struct can then implement the Fly and Crawl methods by delegating to the Bird and Lizard structs.
// The Dragon struct can also implement the Age and SetAge methods by delegating to the Bird and Lizard structs.
// This allows us to create a Dragon object that can fly and crawl, and also has an age.
// The Dragon object is composed of a Bird and a Lizard, and it delegates to them to implement the Fly, Crawl, Age, and SetAge methods.
