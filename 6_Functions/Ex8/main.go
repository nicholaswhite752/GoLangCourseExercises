package main

import "fmt"

func main() {

	x := to10()
	x()
}

func to10() func() {

	f := func() {
		total := 0
		for i := 0; i < 10; i++ {
			total += i
		}
		fmt.Println(total)
	}

	return f
}
