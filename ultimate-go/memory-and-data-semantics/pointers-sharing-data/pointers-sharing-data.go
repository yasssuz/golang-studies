package main

import "fmt"

func main() {
	counter := 10

	//  Value displayed will be 10 and the correct address.
	fmt.Println("Original counter before increment", counter, &counter)

	// We pass the address, not the value
	increment(&counter)

	// Value displayed will be 11, and the address will still be the same.
	// This means that a side effect was created, where down the stack an
	// item in memory of the first stack was modified
	fmt.Println("Original counter after increment", counter, &counter)
}

// `*int` type means that we are expecting the address of an int
func increment(counter *int) {
	*counter++

	// `counter` will be 11, howeven, &counter will be different than the original one.
	// This is because when a parameter is passed into another function, the value is copied.
	// In this case, we are having the copy of the address of the counter
	fmt.Println("Incremented counter:", *counter, &counter)
}
