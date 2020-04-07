package main

import "fmt"

func main() {

	switch {
	case false:
		fmt.Println("This should never print")
	case true:
		fmt.Println("This should always print")
	}

}
