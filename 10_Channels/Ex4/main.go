package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

//Logic
//2 channels q and c
//Put things into c (will be printed in receive function)
//Put 1 into q at end of putting things into c
func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

//Logic
//When c/outPut is done having int put into it, q/done is set
//So a select to print when something in c, and return when something in q
func receive(outPut <-chan int, done <-chan int) {
	for {
		select {
		case v := <-outPut:
			fmt.Println(v)
		case <-done:
			return
		}
	}
}
