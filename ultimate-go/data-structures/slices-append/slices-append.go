package main

import "fmt"

func main() {
	// Declare and initialize a slice to its zero-value
	var s1 []string
	// | nil | -> zero-value in a pointer is nil
	// |  0  | -> length is zero as there is no element
	// |  0  | -> capacity is zero as it was not declared manually so it falls back to the size of the length

	// There is a difference between s1 and s2
	// s1 -> initialized to absolute zero value. The pointer is referencing nil
	// s2 -> initialized to an empty struct. The pointer is pointing to an empty array.
	s2 := []string{}

	fmt.Printf("%#v \n", s1)
	fmt.Printf("%#v \n", s2)

	for record := 1; record <= 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)

		// Append value to the end of the slice
		// We declare a new value of s1 being the copy of s1 (cheap) with appended value
		s1 = append(s1, value)

		// The way append works is:
		// 1. Check if there is enough capacity -> is length the same as capacity?
		// 2. If there is more length than capacity -> allocate the value to the end of the array.
		// 3. If there is no capacity -> double the size of the underlying array.

		// A few caviats here:
		// - Append doesn't change the underlying array, it only returns a copy.
		// - The underlying array will be doubled only until it has 1000 elements, after that is increasing size by 25%.
		// - When creating a new array, the data needs to be copied from the original array to the new array and add the
		//   new appended value on top.
	}
}
