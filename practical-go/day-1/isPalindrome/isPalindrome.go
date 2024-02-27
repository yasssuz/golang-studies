package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(isPalindrome("t"))
	fmt.Println(isPalindrome("to"))
	fmt.Println(isPalindrome("tot"))
	fmt.Println(isPalindrome("toto"))
}

func isPalindrome(text string) bool {
	characters := utf8.RuneCountInString(text)

	return characters%2 != 0
}
