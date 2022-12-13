package main

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	s := Square{}
	s.width = size
	s.height = size
	return &s
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

func UseIt(rc Sized) {
	w := rc.GetWidth()
	rc.SetHeight(10)

	// area := w * rc.height
	area := w * rc.GetHeight()
	println("Expected area of ", 10*w, ", got ", area)
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}

// LSP - Liskov Substitution Principle
// Liskov Substitution Principle (LSP) states that
// "objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program."

// The Liskov Substitution Principle (LSP) is a principle in object-oriented programming and design that states:
// "If for each object o1 of type S there is an object o2 of type T such that for all programs P defined in terms of T,
// the behavior of P is unchanged when o1 is substituted for o2 then S is a subtype of T."

// Rectangle and Square are both subtypes of Sized.
// The LSP states that we should be able to substitute a Rectangle for a Square (or vice versa) and have the program behave correctly.

// In this example, we have a function that calculates the area of a rectangle.
// It takes a Sized object, which is an interface that requires a GetWidth and GetHeight method.
// We can pass in a Rectangle or a Square to this function and it will work correctly.
