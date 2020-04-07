package main

import "fmt"

func main() {

	f := func() {
		total := 0
		for i := 0; i < 10; i++ {
			total += i
		}
		fmt.Println(total)
	}

	f()
	fmt.Printf("%T\n", f)

	g := f
	g()
	fmt.Printf("%T\n", g)
}
