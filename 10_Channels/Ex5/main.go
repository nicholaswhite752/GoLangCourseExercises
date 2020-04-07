package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//Need to assign in go func to avoid deadlock
	//Then can read normally
	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
