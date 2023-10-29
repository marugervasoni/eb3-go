package main

import (
	"fmt"
	"math"
)

//structures
type circle struct {
	radius float64
}
type rect struct {
	width, height float64
}

//geometry interface
type geometry interface {
	area() float64
	perim() float64
}

//constants
const (
	rectType = "RECT"
	circleType = "CIRCLE"
)

func main() {

	c := newGeometry(circleType, 2)
	r := newGeometry(rectType, 2, 3)
	details(c)
	details(r)
	
}

//circle structure methods 
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//details function
// func details(c circle) {
// 	fmt.Println("circle radius: \t",c)
// 	fmt.Println("area: \t\t",c.area())
// 	fmt.Println("perimeter: \t",c.perim())
// }

//rect structure methods 
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//details function modified por geometry
func details(g geometry) {
	fmt.Println("figure geometric: \t",g)
	fmt.Println("area: \t\t\t",g.area())
	fmt.Println("perimeter: \t\t",g.perim())
}

//createGeometry function
func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}	
	}
	return nil
}
