package main

import (
	"fmt"

	"./point"
)

func main() {
	p := point.Point{1, 2}
	q := point.Point{4, 6}
	fmt.Println(point.Distance(p, q))
	fmt.Println(p.Distance(q))
}
