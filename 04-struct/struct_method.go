package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width*r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}


func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())

	c1.setRadius(15)
	fmt.Println("1. radius of c1 is: ", c1.radius)

	c1.setRadius1(16)
        fmt.Println("2. radius of c1 is: ", c1.radius)
}

func (c Circle) setRadius(radius float64) {
	c.radius = radius
}

func (c *Circle) setRadius1(radius float64) {
	c.radius = radius
}

