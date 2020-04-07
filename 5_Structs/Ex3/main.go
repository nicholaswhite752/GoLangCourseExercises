package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Yellow",
		},
		luxury: true,
	}

	fmt.Println("truck1")
	fmt.Printf("\tType: %T\n\tDoors: %v\n\tColor: %v\n\tFourWheel: %v\n", truck1, truck1.doors, truck1.color, truck1.fourWheel)

	fmt.Println("sedan1")
	fmt.Printf("\tType: %T\n\tDoors: %v\n\tColor: %v\n\tFourWheel: %v\n", sedan1, sedan1.doors, sedan1.color, sedan1.luxury)
}
