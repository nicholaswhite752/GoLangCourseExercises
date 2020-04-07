package main

import "fmt"

func main() {

	func() {
		fmt.Println("This is my anon function")
	}()
}
