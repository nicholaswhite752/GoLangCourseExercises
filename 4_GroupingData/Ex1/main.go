package main

import "fmt"

func main() {

	x := [5]int{}
	//BETTER WAY and COMPOSITE LITERAL
	//x := [5]int{0, 100, 200, 300, 400}
	x[0] = 0
	x[1] = 100
	x[2] = 200
	x[3] = 300
	x[4] = 400

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}
