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
	return math.Hypot(q.X-p.X, q.Y-q.Y)
}
