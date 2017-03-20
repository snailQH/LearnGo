package point

import (
	"image/color"
	"math"
)

//Point is an announcement of point type
type Point struct {
	X, Y float64
}

//ColoredPoint is an announcement of color of Point
type ColoredPoint struct {
	Point
	Color color.RGBA
}

//Distance is a traditional function for calculate the distance between two point
func Distance(p, q Point) float64 {
	x1 := (q.X - p.X)
	x2 := (q.Y - p.Y)
	return math.Hypot(x1, x2)
	//return math.Hypot(q.X-p.X, q.Y-q.Y)
}

//Distance is a method of Point type for calculate the distance between two point
func (p Point) Distance(q Point) float64 {
	x1 := (q.X - p.X)
	x2 := (q.Y - p.Y)
	return math.Hypot(x1, x2)
}

//Add is a method for add the value of axis between two point
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

//Sub is a method for sub the value of axis between two point
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

//Path for array of point
type Path []Point

//TranslateBy test
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}
