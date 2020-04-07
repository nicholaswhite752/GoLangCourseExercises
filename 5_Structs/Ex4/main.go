package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
		favs  []string
	}{
		first: "Steve",
		last:  "Smith",
		favs:  []string{"Chocolate", "Vanilla", "Mint"},
	}

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.favs {
		fmt.Printf("\tIndex: %d\tFlavor: %s\n", i, v)
	}

}
