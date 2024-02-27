package main

import "fmt"

func main() {
	passByValue()
}

func passByValue() {
	counter := 10

	// counter refers to the value in memory, &counter refers to the address in memory
	// data is both value and memory
	fmt.Println("Original counter before increment", counter, &counter)

	increment(counter)

	fmt.Println("Original counter after increment", counter, &counter)
}

func increment(counter int) {
	counter++
	fmt.Println("Incremented counter:", counter, &counter)
}
