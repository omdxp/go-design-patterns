package main

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	println("Loading image from", filename)
	return &Bitmap{filename}
}

func (b *Bitmap) Draw() {
	println("Drawing image", b.filename)
}

func DrawImage(image Image) {
	println("About to draw the image")
	image.Draw()
	println("Done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func (b *LazyBitmap) Draw() {
	if b.bitmap == nil {
		b.bitmap = NewBitmap(b.filename)
	}
	b.bitmap.Draw()
}

func main() {
	_ = NewBitmap("photo.jpg") // still loads the image
	// DrawImage(bmp)

	bmp := NewLazyBitmap("photo.jpg") // only loads the image when it's needed for the first time
	DrawImage(bmp)
}

// Virtual Proxy is a structural design pattern that lets you provide a substitute or placeholder for another object.
// A virtual proxy controls access to the original object, allowing you to perform something either remotely or when a resource-intensive object is needed.
