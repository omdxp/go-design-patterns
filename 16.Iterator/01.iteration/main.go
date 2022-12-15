package main

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

type PersonNameIterator struct {
	person *Person
	index  int
}

func (p *Person) Iterator() *PersonNameIterator {
	return &PersonNameIterator{person: p, index: -1}
}

func (i *PersonNameIterator) Next() bool {
	i.index++
	return i.index < 3
}

func (i *PersonNameIterator) Value() string {
	switch i.index {
	case 0:
		return i.person.FirstName
	case 1:
		return i.person.MiddleName
	case 2:
		return i.person.LastName
	default:
		panic("Index out of range")
	}
}

func main() {
	p := Person{"Alexander", "Graham", "Bell"}
	for _, name := range p.Names() {
		println(name)
	}

	for name := range p.NamesGenerator() {
		println(name)
	}

	for i := p.Iterator(); i.Next(); {
		println(i.Value())
	}
}

// Iterator is a behavioral design pattern that lets you traverse elements of a collection without exposing its underlying representation (list, stack, tree, etc.).
