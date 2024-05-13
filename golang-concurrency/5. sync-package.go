package main

import (
	"fmt"
	"sync"
	"time"
)

// The sync package in go contains the concurrency primitives that are most useful for
// low-level memory access synchronization - usually mostly in small scopes such as a struct.

func main() {
	// Wait group is a way to wait for a set of concurrent operation to complete when you either don't care about
	// the result of the concurrent operation, or you have other means of collecting their results.
	// If neither is the case, you should probably skip to channels and the select statement.
	// You can think of WaitGroup like a concurrent-safe counter.
	wg := sync.WaitGroup{}

	// Indicates a delta of the amount of goroutines running.
	// As a rule of thumb, we should only call wg.Add once. We should know how many goroutines we are going to have since
	// the start. We should not create goroutines without knowing when and how the goroutine is going to terminate.
	// It needs to be done outside the goroutine, otherwise you would have a race condition, where the call to Wait could be
	// reached with a zero counter, as the goroutines didn't had enough time to start and signal WaitGroup.
	wg.Add(2)

	fmt.Println("Goroutines will be created")

	go func() {
		// We call wg.Done() before function exits to inform WaitGroup that we've exited
		defer wg.Done()
		fmt.Println("Hello World from Goroutine 1!")
		time.Sleep(1 * time.Second)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Hello World from Goroutine 2!")
		time.Sleep(2 * time.Second)
	}()

	fmt.Println("Goroutines are created")

	// We call wg.Wait() to block the main goroutine until all goroutines have exited.
	wg.Wait()
	fmt.Println("Goroutines have exited")

	mutex()
}

func mutex() {
	// Mutex stands for "mutual exclusion", and is used for memory access synchronization.
	// A mutex provides a concurrent-safe way to express exclusive access to these shared resources.
	var lock sync.Mutex
	var count int
	loopT := 5

	increment := func() {
		lock.Lock()         // Request exclusive access to critical section - in this case count variable.
		defer lock.Unlock() // Before exiting, unlock the critical section.
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var wg sync.WaitGroup
	wg.Add(loopT)
	for i := 0; i < loopT; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Add(loopT)
	for i := 0; i < loopT; i++ {
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines have exited")
}
