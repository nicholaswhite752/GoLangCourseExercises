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

	//This one gives IDE and compiler error
	//saySomething(p1)

	saySomething(&p1)
}

func (p *person) speak() {

	fmt.Printf("My name is %s %s, I'm %d years old", p.first, p.last, p.age)

}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
