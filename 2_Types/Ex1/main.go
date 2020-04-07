package main

import "fmt"

func main() {

	x := 42
	//d is decimal, b is binary, x is hex
	//%x# is hex with the 0x in front
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	fmt.Printf("%d\t%b\t%x\n", x, x, x)

}
