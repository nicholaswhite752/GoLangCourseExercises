package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())
	const gs = 100
	wg.Add(100)

	counter := 0

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			v++
			runtime.Gosched()
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
