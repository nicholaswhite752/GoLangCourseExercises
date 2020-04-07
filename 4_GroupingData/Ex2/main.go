package main

import "fmt"

func main() {

	x := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}
