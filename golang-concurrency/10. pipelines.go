package main

import "fmt"

func main() {
	// A pipeline, in computer science, is nothing more than a series of functions that take data in, perform an
	// operation on it, and pass the data back out - each of those functions are called "stages".
	// A stage always return the same type it consumes. By assuring this, each stage can be modified independently of
	// another, you can mix and match how stages are combined, you can process each stage concurrent to upstream or
	// downstream stages, and even fan-out or rate-limit portions of a pipeline.

	// Batch Processing is when your pipeline is sending to each stage chunks of data.
	// In batch processing, the memory footprint of the program will always be double. This is because a new - in this
	// case, slice - needs to be created within each stage, to not mutate the original data.
	batchProcessing()

	fmt.Println("------------------------------------------------------")

	// Stream Processing is when your pipeline is sending to each stage an individual element at the time.
	// In stream processing, each stage receives and emits a discrete value, and the memory footprint of your program
	// is back down to only the size of the pipeline's input. However, this approach limits the reuse of how the pipeline
	// is fed, and also the pipeline's own ability to scale.
	// Also, you're effectively instantiating your pipeline per each iteration loop... so besides cost, concurrency also
	// becomes an issue.
	streamProcessing()

	fmt.Println("------------------------------------------------------")

	// In this example, we used channels to create our pipeline.
	// There's a lot more code into it, but there's a critical difference: because we are using channels to communicate,
	// each stage of the pipeline can safely execute concurrently - as our inputs and outputs will always be safe
	// within concurrent contexts.
	// Because it executes concurrently, it means that any stage only needs to wait for its inputs, and to be able to
	// send its outputs - aka, it allows stages to execute independent of one another for some slice of time.
	channelsPipeline()
}

func batchProcessing() {
	s := []int{1, 2, 3, 4, 5}
	am := 5 // additional and multiplier

	add := func(vls []int, a int) []int {
		x := make([]int, len(vls))
		for i, n := range vls {
			x[i] = n + a
		}
		return x
	}
	multiply := func(vls []int, m int) []int {
		x := make([]int, len(vls))
		for i, n := range vls {
			x[i] = n * m
		}
		return x
	}

	for _, n := range add(multiply(s, am), am) {
		fmt.Println(n)
	}
}

func streamProcessing() {
	s := []int{1, 2, 3, 4, 5}
	am := 5 // additional and multiplier

	add := func(v int, a int) int {
		return v + a
	}
	multiply := func(v int, m int) int {
		return v * m
	}

	for _, n := range s {
		fmt.Println(add(multiply(n, am), am))
	}
}

func channelsPipeline() {
	generator := func(done <-chan any, ints ...int) <-chan int {
		s := make(chan int, len(ints))
		go func() {
			defer close(s)
			for _, i := range ints {
				select {
				case <-done:
					return
				case s <- i:
				}
			}
		}()
		return s
	}
	add := func(done <-chan any, intStream <-chan int, a int) <-chan int {
		s := make(chan int)
		go func() {
			defer close(s)
			for i := range intStream {
				select {
				case <-done:
					return
				case s <- i + a:
				}
			}
		}()
		return s
	}
	multiply := func(done <-chan any, intStream <-chan int, m int) <-chan int {
		s := make(chan int)
		go func() {
			defer close(s)
			for i := range intStream {
				select {
				case <-done:
					return
				case s <- i * m:
				}
			}
		}()
		return s
	}

	done := make(chan any)
	defer close(done)

	ints := generator(done, 1, 4, 6, 0, 2, 34)
	p := add(done, multiply(done, add(done, ints, 2), 2), 2)
	for i := range p {
		fmt.Println(i)
	}
}
