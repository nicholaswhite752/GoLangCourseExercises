package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())

	wg.Add(2)

	go func() {
		fmt.Println("First Function")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second Function")
		wg.Done()
	}()

	wg.Wait()
}
