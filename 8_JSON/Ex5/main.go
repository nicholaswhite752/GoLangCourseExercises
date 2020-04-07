package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	fmt.Println("------- BY AGE ------")
	sort.Sort(ByAge(users))

	for _, v := range users {
		fmt.Printf("Name: %s %s\n", v.First, v.Last)
		fmt.Printf("\tAge: %d\n", v.Age)
		fmt.Printf("\tSayings:\n")
		sort.Strings(v.Sayings)
		for _, v1 := range v.Sayings {
			fmt.Printf("\t\t%s\n", v1)
		}
		fmt.Println()
	}

	fmt.Println("------- BY Name ------")
	sort.Sort(ByName(users))

	for _, v := range users {
		fmt.Printf("Name: %s %s\n", v.First, v.Last)
		fmt.Printf("\tAge: %d\n", v.Age)
		fmt.Printf("\tSayings:\n")
		sort.Strings(v.Sayings)
		for _, v1 := range v.Sayings {
			fmt.Printf("\t\t%s\n", v1)
		}
		fmt.Println()
	}

}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []user

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Last < a[j].Last }
