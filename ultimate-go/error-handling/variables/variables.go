package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Define base errors
var (
	ErrBadRequest = errors.New("bad request")
	ErrNotFound   = errors.New("page moved")
	ErrInternal   = errors.New("internal server error")
)

func main() {
	if such, err := webCall(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(such)
	}
}

func webCall() (string, error) {
	// Return different values based on the generated number
	randNum := rand.Intn(3) + 1 // rand.Intn will return between 0 and 2, so we add a number to make it between 1 and 3

	switch randNum {
	case 1:
		return "", ErrBadRequest
	case 2:
		return "", ErrNotFound
	case 3:
		return "Success", nil
	default:
		return "", ErrInternal
	}
}
