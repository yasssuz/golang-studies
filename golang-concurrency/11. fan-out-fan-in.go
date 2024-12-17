package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Sometimes, stages in your pipeline can be particularly computationally expensive. When this happens, upstream
	// stages in your pipeline can become blocked while waiting for your expensive stages to complete. Not only that,
	// but the pipeline itself can take a long time to execute as a whole.

	// One solution to this is using the Fan-Out Fan-In approach, which consists of reusing a single stage of the
	// pipeline on multiple goroutines in an attempt to parallelize pulls from an upstream stage.
	// Fan-Out is a term to describe the process of starting multiple concurrent processes to handle input from a given
	// pipeline, and Fan-In is a term to describe the process of combining multiple results into one channel.

	// You might consider fanning out one of your stages if both of the following apply:
	//		- It doesn't rely on values that the stage had calculated before.
	//		- It takes a long time to run.

	start := time.Now()

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pipeline := stage4(stage3(stage2(stage1(data))))
	for i := range pipeline {
		fmt.Println(i)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}

// Stage 1: Generate data
func stage1(data []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, d := range data {
			out <- d
		}
		close(out)
	}()
	return out
}

// Stage 2: Process data
func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for d := range in {
			// Perform some processing
			out <- d + 1
		}
		close(out)
	}()
	return out
}

// Stage 3: Time-consuming task (to be implemented with fan-out fan-in)
func stage3(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// Simulate time-consuming processing
		for d := range in {
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond) // Simulate delay
			out <- d * 2                                                  // Dummy processing: double the value
		}
		close(out)
	}()
	return out
}

// Stage 4: Finalize and output result
func stage4(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for d := range in {
			// Final processing
			out <- d * 2
		}
		close(out)
	}()
	return out
}
