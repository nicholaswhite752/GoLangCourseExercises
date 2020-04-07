package main

import "fmt"

const (
	year = 2020 + iota
	a    = year + iota
	b    = year + iota
	c    = year + iota
	d    = year + iota
)

func main() {

	fmt.Println(year, a, b, c, d)

}
