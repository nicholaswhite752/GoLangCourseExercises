package main

import "fmt"

//Creating my own type
type myCreation int

//Creating a variable named x, of type myCreation
var x myCreation

func main() {

	//prints value of x
	//could also do comment below for just value
	//fmt.Println(x)
	fmt.Printf("%v\n", x)
	//prints type of x
	fmt.Printf("%T\n", x)

	//assign x to 42
	x = 42
	//prints value of x
	fmt.Printf("%v\n", x)

}
