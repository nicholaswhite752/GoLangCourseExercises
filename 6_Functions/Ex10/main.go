package main

import "fmt"

func main() {

	a := countUp()
	b := countUp()

	for i := 0; i < 15; i++ {
		a()
	}

	for i := 0; i < 25; i++ {
		b()
	}

	fmt.Println(a())
	fmt.Println(b())

}

func countUp() func() int {
	var x int
	return func() int {
		x++
		return x
	}

}
