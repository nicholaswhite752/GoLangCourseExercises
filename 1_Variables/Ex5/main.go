package main

import "fmt"

//Creating my own type
type myCreation int

//Creating a variable named x, of type myCreation
var x myCreation

//create var y of underlying type int
var y int

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

	//Example of Conversion
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
