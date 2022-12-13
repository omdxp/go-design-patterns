package main

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level module
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module
type Research struct {
	// break DIP by depending on a low-level module
	relationships Relationships
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			println("John has a child called", rel.to.name)
		}
	}
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships2 struct {
	relations []Info
}

func (r *Relationships2) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.from.name == name && v.relationship == Parent {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Research) Investigate2(browser RelationshipBrowser) {
	for _, p := range browser.FindAllChildrenOf("John") {
		println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	// high-level module
	research := Research{relationships}
	research.Investigate()

	// DIP
	relationships2 := Relationships2{}
	relationships2.relations = append(relationships2.relations, Info{&parent, Parent, &child1})
	relationships2.relations = append(relationships2.relations, Info{&parent, Parent, &child2})
	research.Investigate2(&relationships2)
}

// DIP - Dependency Inversion Principle
// Dependency Inversion Principle (DIP) states that high-level modules should not depend on low-level modules. Both should depend on abstractions.

// Abstractions should not depend on details. Details should depend on abstractions.
