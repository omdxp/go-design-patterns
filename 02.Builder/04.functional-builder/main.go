package main

import "fmt"

type Person struct {
	name, position string
}

type PersonMod func(*Person)
type PersonBuilder struct {
	actions []PersonMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func main() {
	pb := PersonBuilder{}
	pb.Called("Omar").Called("Belghaouti").WorksAsA("Software Engineer")
	p := pb.Build()
	fmt.Println(*p)
}

// functional builder is a pattern that allows us to build objects using a fluent interface.
// The difference between this and the builder parameter is that the builder parameter
// is a function that takes a builder as a parameter, whereas the functional builder
// is a function that returns a builder.
