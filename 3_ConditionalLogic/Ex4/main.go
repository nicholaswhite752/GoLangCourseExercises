package main

import "fmt"

func main() {

	i := 1996

	for {
		if i > 2020 {
			break
		}

		fmt.Printf("I was alive in %v\n", i)
		i++
	}

}
