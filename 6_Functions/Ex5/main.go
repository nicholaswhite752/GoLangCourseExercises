package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func main() {

	sq1 := square{length: 10}
	cir1 := circle{radius: 10}

	info(sq1)
	info(cir1)

}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println("The shape's area is: ", sh.area())
}

func (s square) area() float64 {

	return s.length * s.length

}

func (c circle) area() float64 {

	return c.radius * c.radius * math.Pi

}
