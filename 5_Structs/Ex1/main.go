package main

import "fmt"

type person struct {
	first string
	last  string
	favs  []string
}

func main() {

	p1 := person{
		first: "Steve",
		last:  "Smith",
		favs:  []string{"Chocolate", "Vanilla", "Mint"},
	}

	p2 := person{
		first: "Jane",
		last:  "Doe",
		favs:  []string{"Strawberry", "Oreo", "Cotton Candy"},
	}

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.favs {
		fmt.Printf("\tIndex: %d\tFlavor: %s\n", i, v)
	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.favs {
		fmt.Printf("\tIndex: %d\tFlavor: %s\n", i, v)
	}

}
