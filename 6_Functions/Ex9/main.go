package main

import "fmt"

func main() {
	//Takes in positive number and adds all numbers before it together and makes negative
	//Convoluted way, but shows function calling a function
	xNeg := addNeg(toNeg, 10)
	fmt.Println(xNeg)
}

func toNeg(ni int) int {

	return ni * -1

}

func addNeg(f func(int) int, number int) int {
	total := 0

	for i := 0; i < number; i++ {
		temp := f(i)
		total += temp
	}

	return total
}
