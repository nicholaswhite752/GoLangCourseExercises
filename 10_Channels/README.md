===========================

Ex1)

● get this code working:
	○ using func literal, aka, anonymous self-executing func 
	○ a buffered channel


```
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}

```

===========================

Ex2)

● Get this code running:

```
package main

import (
	"fmt"
)

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

```

```
package main

import (
	"fmt"
)

func main() {
	cr := make(<-chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

```



===========================

Ex3)

Starting with this code, pull the values off the channel using a for range loop

```
package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}

```


===========================

Ex4)

Starting with this code, pull the values off the channel using a select statement

```
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}

```


===========================

Ex5)

Show the comma ok idiom starting with this code.

```
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	v, ok :=
	fmt.Println(v, ok)

	close(c)
	
	v, ok =
	fmt.Println(v, ok)
}

```

===========================

Ex6)

● write a program that
	○ puts 100 numbers to a channel
	○ pull the numbers off the channel and print them


===========================

Ex7)

● write a program that
	○ launches 10 goroutines
		■ each goroutine adds 10 numbers to a channel
	○ pull the numbers off the channel and print them


===========================


