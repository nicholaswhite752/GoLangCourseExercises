package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "Steve",
		last:  "Smith",
		age:   35,
	}

	p1.speak()
}

func (p person) speak() {

	fmt.Printf("My name is %s %s, I'm %d years old", p.first, p.last, p.age)

}
