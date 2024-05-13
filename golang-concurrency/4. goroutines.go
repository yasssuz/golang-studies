package main

import (
	"fmt"
	"sync"
)

// In go, a "goroutine" is simply a branch-of-execution (or a function) that is running concurrently to the main branch
// of execution. Remember, concurrent means it runs out of order.

// Under the hood, Goroutines are "coroutines", which is a high level abstraction that leverages Green Threads. Coroutines
// are non-preemptive - they cannot be interrupted. Instead, they have multiple points throughout which allow for
// suspension or reentry. They don't define their own suspension or reentry points, instead Go's runtime observes the
// runtime behavior and automatically suspense them when they block and resumes when they become unblocked.

func main() {
	// sync.WaitGroup helps us maintain a count of how many goroutines exist an any given time
	var wg sync.WaitGroup

	// Here we say that we are going to have 2 goroutines
	// As a rule of thumb, we should only call wg.Add once. We should know how many goroutines we are going to have since
	// the start. We cannot create goroutines without knowing when and how the goroutine is going to terminate.
	wg.Add(2)

	fmt.Println("Goroutines starting")

	// Create goroutine to run function lowercase
	go func() { // Fork-point 1
		lowercase()
		defer wg.Done()
	}()

	// Create goroutine to run function uppercase
	go func() { // Fork-point 2
		uppercase()
		defer wg.Done()
	}()

	fmt.Println("Goroutines waiting to finish")

	// Blocks the main goroutine until all the other goroutines have terminated
	wg.Wait() // Join-point

	fmt.Println("Goroutines finished")
}

func lowercase() {
	fmt.Println("lowercase")
	fmt.Println("lowercase")
	fmt.Println("lowercase")
	fmt.Println("lowercase")
	fmt.Println("lowercase")
	fmt.Println("lowercase")
}

func uppercase() {
	fmt.Println("uppercase")
	fmt.Println("uppercase")
	fmt.Println("uppercase")
	fmt.Println("uppercase")
	fmt.Println("uppercase")
	fmt.Println("uppercase")
}
