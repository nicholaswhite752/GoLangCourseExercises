package main

import "fmt"

func main() {

	//Those characters ` are the thing on the ~/tilde key, for string literals
	//It is not the '/apostrophe, those are for rune literals
	x := `"Here is a string literal"`

	fmt.Println(x)
	fmt.Println("")

	//here is another good example with quotes and returns
	a := `here is something
as 
a 
raw string
literal
"you see"
another thing`
	fmt.Println(a)

}
