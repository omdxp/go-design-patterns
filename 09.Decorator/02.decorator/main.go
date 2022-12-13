package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("A circle of radius %.2f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("A square with side %.2f", s.Side)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %.2f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

func main() {
	circle := Circle{2}
	circle.Resize(2)
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "red"}
	fmt.Println(redCircle.Render())

	redHalfCircle := ColoredShape{&Circle{5}, "red"}
	fmt.Println(redHalfCircle.Render())

	// redHalfCircle.Resize(2) // error: cannot use redHalfCircle (type ColoredShape) as type Circle in argument to redHalfCircle.Resize:
	// ColoredShape does not implement Circle (missing Resize method)

	transRedHalfCircle := TransparentShape{&redHalfCircle, 0.5}
	fmt.Println(transRedHalfCircle.Render())
}

// Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.

// The Decorator pattern is often useful for adhering to the Single Responsibility Principle.
// It lets you divide a monolithic class that implements many possible variants of behavior into several separate classes.
