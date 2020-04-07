package main

import "fmt"

func main() {
	c := make(chan int)

	//Need to assign in go func to avoid deadlock
	//Then can read normally
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Leaving Program")

}
