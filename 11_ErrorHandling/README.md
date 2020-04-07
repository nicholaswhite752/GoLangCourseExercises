===========================

Ex1)

Start with this code. Instead of using the blank identifier, make sure the code is checking and handling the error.

```
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))

}


```

===========================

Ex2)

Start with this code. Create a custom error message using “fmt.Errorf”.

```
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) []byte {
	bs, _ := json.Marshal(a)
}

```



===========================

Ex3)

Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. 


===========================

Ex4)

Starting with this code, use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your 
● lat "50.2289 N"
● long "99.4656 W"


```
package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
	}
	return 42, nil
}


```


===========================




