package main

import "fmt"

func main() {
	// Decleare and initialize a slice of length 5 with zero-value.
	fruits := make([]string, 5)

	// The main difference between arrays and slices is:
	// 1. Arrays are fixed size, while slices are not.
	// 2. A slice is a three words data structure:
	// | 0xn | -> pointing to backing array
	// | len | -> total number of elements in the slice from 0 to x that we can read or write to.
	// | cap | -> represents the total number of elements that exist in the backing array.

	// Easiest way to differentiate an array is by looking at the position of the length and the function.
	// To initialize a slice at zero-value you need to leverage the make function, is not the case with an array.
	// Also,
	// Array (zero-value) -> var items [2]string
	// Array (with value) -> items := [2]string{"Apple", "Pear"}
	// Slice (zero-value) -> items := make([]string, 2)
	// Slice (with value) -> items := []string{"Apple", "Pear"}

	// Can also be declared as
	// fruits := []string{"Apple", "Pear", "Kiwi", "Banana", "Watermelon"}

	// Can assign values on index like an array
	fruits[0] = "Apple"
	fruits[1] = "Pear"
	fruits[2] = "Kiwi"
	fruits[3] = "Banana"
	fruits[4] = "Watermelon"

	// panic: "runtime error: index out of range [5] with length 5"
	// fruits[5] = "Grapes"

	// The length of the array is 5, you cannot change value on a index that doesn't exist.
	// The only way you can do this is by appending, which would enforce the slice to double its size and
	// length would become 10.

	// Decleare and initialize a slice of length 5 with zero-value, and set its capacity to 8.
	people := make([]string, 5, 8)

	// panic: "runtime error: index out of range [5] with length 5"
	// people[5] = "Karim"

	// This will result in error even tho the capacity is 8.
	// Remember, length is the total amount of elements we can read or write from zero, and it is still 5.
	// Setting capacity to 8 just helps with future effeciency for growth.

	inspectSlice(people)
}

// Slices should be passed as value, not by reference.
// This is because the underlying backing array (heavy part) is already passed by reference, and
// the other two words (remember, a slice has 3 words) are cheap to copy.
func inspectSlice(slice []string) {
	fmt.Printf("Length: %d, Capacity: %d", len(slice), cap(slice))

	// Note that we are using value semantics for this for range loop.
	// What is this slice made of? -> strings
	// What is a string? -> built-in type
	// built-in types are should always use value semantics
	for i, v := range slice {
		fmt.Print(i, v)
	}
}
