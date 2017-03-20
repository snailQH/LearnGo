package main

import (
	"fmt"
	"math"
)

//Point define the new type
type Point struct{ X, Y float64 }

//Distance ,traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.Y, q.Y-p.Y)
}

//Distance ,same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) //print 4.47213595499958
	fmt.Println(p.Distance(q))  //print 5
}
