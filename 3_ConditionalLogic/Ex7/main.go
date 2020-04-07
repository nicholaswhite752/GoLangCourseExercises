package main

import "fmt"

func main() {

	x := 50

	if x < 10 {
		fmt.Println("X is less than 10")
	} else if x < 75 {
		fmt.Println("X is less than 75")
	} else {
		fmt.Println("X is greater than or equal to 100")
	}

}
