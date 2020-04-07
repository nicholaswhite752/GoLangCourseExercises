package main

import "fmt"

//Default values are called ZERO values
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	//I wrote this
	//s := fmt.Sprintln(x, y, z)

	//Could also do below, which gives tabs and format
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

}
