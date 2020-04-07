package main

import "fmt"

type custErr struct {
	s string
}

func main() {

	cust1 := custErr{s: "MY CUSTOM ERROR"}
	foo(cust1)

}

func (ce custErr) Error() string {
	return ce.s
}

func foo(e error) {
	fmt.Println(e)
}
