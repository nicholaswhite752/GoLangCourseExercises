package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i <= 9; i++ {
		go func(t int) {
			for j := 0; j <= 9; j++ {
				c <- 10*t + j
			}
			wg.Done()
		}(i)
	}

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	wg.Wait()
	fmt.Println("ROUTINES", runtime.NumGoroutine())
	fmt.Println("Leaving Program")

}
