package main

import "fmt"

func main() {

	foo()
}

func foo() {
	defer func() {
		fmt.Println("This ran after a defer")
	}()

	fmt.Println("This ran at the end of the foo function")
}
