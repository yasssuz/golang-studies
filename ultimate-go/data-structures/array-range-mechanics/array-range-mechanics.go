package main

import "fmt"

func main() {
	friends1 := [5]string{"Karim", "Madalina", "Samir", "Nadia", "Miranda"}

	fmt.Printf("friends1 (Madalina) -> %#v \n", friends1[1])
	for i := range friends1 {
		friends1[1] = "Buni"

		if i == 1 {
			fmt.Printf("friends1 (Madalina) -> %#v \n", friends1[1])
		}
	}

	fmt.Println("**************************************************")

	friends2 := [5]string{"Karim", "Madalina", "Samir", "Nadia", "Miranda"}

	fmt.Printf("frieds2 (Madalina) -> %#v \n", friends2[1])

	// Just by adding `v`, you're iterating over the copy of the array.
	// So even if frieds2 change after one iteration, is not going to reflect on the overall iteration
	for i, v := range friends2 {
		friends2[1] = "Buni"

		if i == 1 {
			fmt.Printf("frieds2 (Madalina) -> %#v", v)
		}
	}
}
