package main

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	// ...
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for _, p := range products {
		if p.color == color {
			result = append(result, &p)
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for _, p := range products {
		if p.size == size {
			result = append(result, &p)
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for _, p := range products {
		if p.size == size && p.color == color {
			result = append(result, &p)
		}
	}
	return result
}

// Specification pattern
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type AndSpecification struct {
	first, second Specification
}

func (a *AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct {
	// ...
}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for _, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &p)
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", red, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	f := Filter{}
	greenThings := f.FilterByColor(products, green)
	for _, v := range greenThings {
		println(v.name, "is green")
	}

	bf := BetterFilter{}
	greenSpec := ColorSpecification{green}
	greenThings = bf.Filter(products, &greenSpec)
	for _, v := range greenThings {
		println(v.name, "is green")
	}
}

// OCP - Open-Closed Principle
// open for extension, closed for modification

// Filter has a violation of OCP
// because if we want to add a new filter, we need to modify the Filter struct
// and add a new method to the Filter struct

// BetterFilter is a solution to the OCP violation
// we can add a new filter by adding a new struct that implements the Specification interface
