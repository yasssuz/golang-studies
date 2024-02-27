package main

import "fmt"

func main() {
	banner("Hey Madalina!", 5)
}

func banner(text string, width int) {
	padding := (width - len(text)) / 2

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}
