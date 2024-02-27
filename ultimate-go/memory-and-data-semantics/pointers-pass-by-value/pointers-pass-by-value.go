package main

import "fmt"

func main() {
	counter := 10

	//  Value displayed will be 10 and the correct address.
	fmt.Println("Original counter before increment", counter, &counter)

	// We pass the value
	increment(counter)

	// Value displayed will still be 10, and the address will still be the same.
	// This means that changing the parameter does not reflect on the original value, preserving immutability.
	fmt.Println("Original counter after increment", counter, &counter)
}

func increment(counter int) {
	counter++

	// `counter` will be 11, howeven, &counter will be different than the original one.
	// This is because when a parameter is passed into another function, the value is copied.
	fmt.Println("Incremented counter:", counter, &counter)
}
