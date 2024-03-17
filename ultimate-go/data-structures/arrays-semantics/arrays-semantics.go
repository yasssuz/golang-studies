package main

import "fmt"

func main() {

	// Declare an array of 5 values that is initialized to its zero value
	var fruits [5]string
	// Below is the demonstration of what was created. You may notice two rows, is because of how strings are in Go.
	// | nil | nil | nil | nil | nil |
	// | --- | --- | --- | --- | --- |
	// | 0   | 0   | 0   | 0   | 0   |

	fmt.Printf("fruits (type) -> %#v \n", fruits)
	fmt.Printf("fruits (value) -> %d \n", fruits)

	fruits[0] = "Apple"
	fruits[1] = "Pear"
	fruits[2] = "Kiwi"
	fruits[3] = "Banana"
	fruits[4] = "Watermelon"

	// | 0xn | 0xy | 0xu | 0xt | 0xb |
	// | --- | --- | --- | --- | --- |
	fmt.Println("************AFTER ASSIGNMENT************")

	fmt.Printf("fruits (type) -> %#v \n", fruits)
	fmt.Printf("fruits (value) -> %d \n", fruits)

	// Iterate over fruits account
	// `fruit` is a copy of fruits[i], however, is very efficient as its pointing to
	// its underlying bytes array -> [0xn \ 5]. Note that this is only with strings.
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}
