package main

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	println("Hello, my name is", p.name, "and I am tired")
}

func (p *person) SayHello() {
	println("Hello, my name is", p.name)
}

func NewPerson(name string, age int) Person {
	if age < 0 {
		age = 0
	}
	if age > 100 {
		return &tiredPerson{name: name, age: age}
	}
	return &person{name: name, age: age}
}

func main() {
	p := NewPerson("John", 20)
	p.SayHello()

	p = NewPerson("John", 200)
	p.SayHello()
}

// interface factory is a factory that returns an interface
// instead of a concrete type
// it is used to hide the implementation details of the struct
// it is used to create a new instance of a struct with some default values
