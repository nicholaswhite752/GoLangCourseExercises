package main

import "fmt"

func main() {

	favSport := "Baseball"

	switch favSport {
	case "Football":
		fmt.Println("My favorite sport is football")
	case "Baseball":
		fmt.Println("My favorite sport is baseball")
	default:
		fmt.Println("My favorite sport is not supported")
	}

}
