package main

import "fmt"

func main() {
	s1 := make([]string, 5, 8)

	s1[0] = "Apple"
	s1[1] = "Pear"
	s1[2] = "Kiwi"
	s1[3] = "Banana"
	s1[4] = "Watermelon"

	fmt.Printf("Length[%d] Capacity[%d] \n", len(s1), cap(s1))

	for i := range s1 {
		v := s1[i]
		fmt.Printf("%d %v %v \n", i, &v, v)
	}

	fmt.Println("************************************************")

	// slicing operation -> take from the 2nd item (3rd index) until but not including the 4th item (5th index)
	s2 := s1[2:4] // -> {"Kiwi", "Banana"}

	fmt.Printf("Length[%d] Capacity[%d] \n", len(s2), cap(s2))

	for i := range s2 {
		v := s2[i]
		fmt.Printf("%d %v %v \n", i, &v, v)
	}

	fmt.Println("************************************************")

	s2[1] = "CHANGED"

	fmt.Printf("S1 -> Length[%d] Capacity[%d] \n", len(s1), cap(s1))
	fmt.Printf("S2 -> Length[%d] Capacity[%d] \n", len(s2), cap(s2))

	fmt.Println("S1 ->")
	for i := range s1 {
		v := s1[i]
		fmt.Printf("%d %v %v \n", i, &v, v)
	}

	fmt.Println("S2	 ->")
	for i := range s2 {
		v := s2[i]
		fmt.Printf("%d %v %v \n", i, &v, v)
	}

	// As you can see, both changed.
	// This is because you have the same underlying backing array as reference, for both slices.
}
