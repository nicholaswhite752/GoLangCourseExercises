package main

import "fmt"

func main() {

	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["smith_steve"] = []string{"Beers", "Beaches", "Protein"}

	delete(x, "no_dr")

	for i0, v0 := range x {
		fmt.Printf("Index: %s\t Contents: %s\n", i0, v0)
		for i1, v1 := range v0 {
			fmt.Printf("\tIndex: %d\t Contents: %s\n", i1, v1)
		}
	}
}
