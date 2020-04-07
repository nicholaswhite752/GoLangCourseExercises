package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "John",
		last:  "Smith",
		age:   35,
	}

	fmt.Println(p1)

	changeMe(&p1, "Mike")

	fmt.Println(p1)

}

func changeMe(pChange *person, changeTo string) {
	pChange.first = changeTo
	//ALSO WORKS
	//(*pChange).first = changeTo
}
