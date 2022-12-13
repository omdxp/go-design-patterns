package main

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float32) {
	println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	vector := VectorRenderer{}
	circle := NewCircle(&vector, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()

	raster := RasterRenderer{}
	circle = NewCircle(&raster, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
}

// Bridge is a structural design pattern that lets you split a large class or a set of closely related classes into two separate hierarchies—abstraction and implementation—which can be developed independently of each other.
