package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	//Need to make go func to run concurrent
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

//prints out channel until close
func receive(out <-chan int) {
	for v := range out {
		fmt.Println(v)
	}
}
