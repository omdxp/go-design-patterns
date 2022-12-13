package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

// given interface
func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}

// interface we have
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(image RasterImage) string {
	maxX, maxY := 0, 0
	for _, p := range image.GetPoints() {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	maxX += 1
	maxY += 1
	imageData := make([][]rune, maxY)
	for i := range imageData {
		imageData[i] = make([]rune, maxX)
		for j := range imageData[i] {
			imageData[i][j] = ' '
		}
	}
	for _, p := range image.GetPoints() {
		imageData[p.Y][p.X] = '*'
	}
	var sb strings.Builder
	for _, row := range imageData {
		sb.WriteString(string(row))
		sb.WriteRune('\n')
	}
	return sb.String()
}

// Adapter
type vectorToRasterAdapter struct {
	points []Point
}

var pointCache = map[[16]byte][]Point{}

func (v *vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func (v *vectorToRasterAdapter) addLine(line Line) {
	left := line.X1
	right := line.X2
	top := line.Y1
	bottom := line.Y2
	if left > right {
		left, right = right, left
	}
	if top > bottom {
		top, bottom = bottom, top
	}
	dx := right - left
	dy := bottom - top
	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}
}

func (v *vectorToRasterAdapter) addLineCached(line Line) {
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}

	h := hash(line)
	if points, ok := pointCache[h]; ok {
		v.points = append(v.points, points...)
		return
	}

	left := line.X1
	right := line.X2
	top := line.Y1
	bottom := line.Y2
	if left > right {
		left, right = right, left
	}
	if top > bottom {
		top, bottom = bottom, top
	}
	dx := right - left
	dy := bottom - top
	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	}
	if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}
	pointCache[h] = v.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := &vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLineCached(line)
	}
	return adapter
}

func main() {
	rc := NewRectangle(10, 5)
	a := VectorToRaster(rc)
	_ = VectorToRaster(rc) // 2nd adapter
	fmt.Print(DrawPoints(a))
}

// Adapter caching (memoization) is a technique that can be used to improve the performance of an adapter.
// The idea is to cache the results of the adapter so that if the same request comes in again, the cached result is returned instead of recalculating it.
// This is a simple example of the memoization pattern.
