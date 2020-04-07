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

	x := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for i0, v0 := range x {
		fmt.Printf("Index: %s\n", i0)
		fmt.Printf("\tFirst\\Last: %s %s\n", v0.first, v0.last)
		fmt.Printf("\tFav Flavors:\n")
		for i1, v1 := range v0.favs {
			fmt.Printf("\t\tIndex: %d\t Contents: %s\n", i1, v1)
		}
	}

}
