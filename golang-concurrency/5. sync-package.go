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
	fmt.Println("**************************************************************************************************")

	//mutex()
	//condSignal()
	//condBroadcast()
	//once()
	poolEx1()
}

func mutex() {
	// Mutex stands for "mutual exclusion", and is used for memory access synchronization.
	// A mutex provides a concurrent-safe way to express exclusive access to these shared resources.
	var lock sync.Mutex

	// A critical section is a segment of code within a program that must be executed atomically, meaning it should
	// not be interrupted or concurrently accessed by multiple threads or goroutines. Critical sections are typically
	// used when working with shared resources, such as variables or data structures, that could be accessed or modified
	// by multiple concurrent execution contexts. When working with critical sections, we usually want to:
	// 	- Keep them small and focused, if multiple operations are needed, consider breaking them up
	// 	- Exit them as fast as possible, avoid I/O blocking or long-running tasks
	//  - Use fina-grained locks, instead of locking an entire data structure, consider locking separate parts of it.
	//  - Avoid nested locks, they increase the likelihood of causing deadlocks
	//  - Consider RWMutex, when there's a lot of reads and little writes.

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

func rwMutex() {
	// TODO write example
}

func condSignal() {
	var wg sync.WaitGroup
	var m sync.Mutex

	// Cond is a rendezvous point for goroutines waiting for or announcing the occurrence of an event.
	// An event, in Go, is any arbitrary signal between two or more goroutines that carries no information other than
	// the fact it has occurred.
	c := sync.NewCond(&m)

	// Very often, you will want wait for one of these signals before continuing execution on a goroutine.
	// One naive approach would be
	// if conditionTrue() == false {
	//	 time.Sleep(1 * time.Millisecond)
	// }
	// One problem with this approach is that it is either CPU or time inefficient. Look, if the waiting time is too
	// long, you're artificially degrading performance, and if it is too small, you're consuming too much CPU time.
	// This is what the cond type fixes.

	// mention how it blocks and suspends go routines - freeing the OS thread for another green thread (aka goroutine)

	// Create a queue. We are going to push 10 elements, 2 at the time.
	queue := make([]any, 0, 10)

	removeFromQueue := func(t time.Duration) {
		time.Sleep(t) // Artificially simulate blocking period

		c.L.Lock()         // Enter critical section
		defer c.L.Unlock() // Exit critical section

		queue = queue[1:] // Remove head of queue
		fmt.Println("Removed from queue")
		c.Signal() // Let the LONGEST WAITING goroutine know that something happened.
		wg.Done()  // Tell wg that the process in this goroutine has finalized
	}

	for i := 0; i < 10; i++ {
		c.L.Lock() // Enter critical section
		wg.Add(1)  // Tell wg that a goroutine was created
		for len(queue) == 2 {
			// Block and suspends main goroutine until a new signal occurs.
			// This is to free the OS Thread and to allow another green thread to take its place.
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})    // Append item to queue
		go removeFromQueue(time.Millisecond) // Remove item from queue after 1 millisecond
		c.L.Unlock()
	}

	// Wait for remaining items in the queue to finish executing
	wg.Wait()
	fmt.Println("Goroutines have exited")
	fmt.Println("**************************************************************************************************")
}

func condBroadcast() {
	type Button struct {
		Clicked *sync.Cond
	}

	// Encapsulate rendezvous point for btn
	btn := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()

			c.L.Lock()
			defer c.L.Unlock()

			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	// Create a wg for button click
	var btnClickWg sync.WaitGroup
	btnClickWg.Add(3) // Add the 3 subscriptions

	// Create subscribed functions (they will run on their own goroutine)
	subscribe(btn.Clicked, func() {
		fmt.Println("Window maximized")
		btnClickWg.Done() // Tell to wg when goroutine finished executing
	})
	subscribe(btn.Clicked, func() {
		fmt.Println("Annoying popup shows up")
		btnClickWg.Done()
	})
	subscribe(btn.Clicked, func() {
		fmt.Println("Colors have changed")
		btnClickWg.Done()
	})

	btn.Clicked.Broadcast() // Broadcast event to all running goroutines
	btnClickWg.Wait()       // Wait for all events to complete
}

func once() {
	var o sync.Once
	// sync.Once ensures that only one call to Do ever call the function passed in.
	// This keeps being true even on different goroutines, and even with different functions.

	var wg sync.WaitGroup
	var l sync.Mutex
	var count int

	increment := func() {
		l.Lock()
		defer l.Unlock()

		count++
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			o.Do(increment)
		}()
	}

	fmt.Println(count) // will display 1 instead of 10
}

func poolEx1() {
	// Pool is a concurrent-safe implementation of the object pool pattern.
	// At a high level, the pool pattern is a way to create and make available a fixed number, or pool, of things for use.
	// It's commonly used to constrain the creation of things that are expensive (e.g. database connections), so that
	// only a fixed number of them are ever created, but an indeterminate number of operations can still request
	// access to these things.

	// TODO write example
}
