package main

import (
	"fmt"
	"sort"
)

func main() {

	// Slices are built on top of arrays
	// array -> fixed size, pass-by-value (copy)
	// slice -> dynamic size, pass-by-reference (pointer)

	// Slices are "nil safe" -> checking nil to an empty slice will result in truthy
	var s1 []int // s is a slice of int
	fmt.Printf("s1 = %#v\n", s1)

	// Each slice has 3 core values on it.
	// Array -> pointer to the underlying array
	// Len (length) -> number of elements it contains
	// Cap (capacity) -> maximun number of elements an array can hold without requiring additional memory

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8} // initializing a slice with some values
	// The above slice will have a cap of 8 and a length of 8
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:3] // slicing from index zero to index 3, which means, get index 1 and 2
	// The above slice will have a cap of 7 and a length of 2
	fmt.Printf("s3 = %#v\n", s3)

	s3 = append(s3, 100) // re-declare s3 with new appended value, mutability
	fmt.Printf("s3 (append) = %#v\n", s3)
	fmt.Printf("s2 (after s3 append) = %#v\n", s2) // also changed

	// s2 changed because slices are passed by reference, not by value.
	// This means that, s3 is referencing to the same underlying array as s2, which
	// means that, if s3 is modified, s2 will reflect those changes.

	// This will panic as is out of range, aka there is nothing till range 9
	// fmt.Print(s2[:9])

	// This will not panic as you're taking from the first to the last value
	fmt.Println(s2[:])

	// Will print both slices merged
	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"}))

}

func concat(s1, s2 []string) []string {
	// s2... means that we are destructuring the slice
	return append(s1, s2...)
}

func median(values []float64) (float64, error) {

	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	// As previously mentioned, when passing slices to other functions, they are passed by reference, not by value.
	// If we sort `values`, we are gonna sort the original values of the underlying array in memory.
	// Is important to copy so that you don't create any unwanted side-effects.

	nums := make([]float64, len(values))
	copy(values, nums)

	sort.Float64s(nums)
	i := len(nums) / 2

	if len(nums)%2 == 1 {
		return nums[i], nil
	}

	v := (nums[i-1] + nums[i]) / 2
	return v, nil
}
