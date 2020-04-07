package main

import (
	"fmt"
)

func main() {
	//Started with
	//cs := make(chan<- int)
	//Changed to
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
