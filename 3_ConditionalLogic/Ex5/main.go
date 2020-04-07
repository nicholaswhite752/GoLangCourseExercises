package main

import "fmt"

func main() {
	//%% is the escape for %
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v%%4 is:\t%v\n", i, i%4)
	}

}
