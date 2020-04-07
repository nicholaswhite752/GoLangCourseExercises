package main

import "fmt"

func main() {

	temp1 := []string{"James", "Bond", "Shaken, not stirred"}
	temp2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(temp1)
	fmt.Println(temp2)

	totalSlice := [][]string{temp1, temp2}
	fmt.Println(totalSlice)

	for i0, v0 := range totalSlice {
		fmt.Printf("Record: %d\t Contents: %v\n", i0, v0)
		for i1, v1 := range v0 {
			fmt.Printf("\tIndex: %d\t Contents: %v\n", i1, v1)
		}

	}

}
