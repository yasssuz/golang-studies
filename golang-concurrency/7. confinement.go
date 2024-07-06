package main

import "fmt"

// When working with concurrent code, there are a few different options for safe operations, like:
// 	- Synchronization primitives for haring memory (e.g., Mutexes)
// 	- Synchronization via communicating (e.g., channels)
// However, there are other options that are implicitly safe within multiple concurrent processes:
// 	- Immutable data -> each concurrent operation may only read on the data or change a copy of it.
// 	- Data protected by confinement -> ensuring data is only ever available from one concurrent process.

// There are two types of confinements:
// 	- Adhoc confinement -> achieving confinement through a convention set by the community or the project
//  - Lexical confinement -> using the lexical scope to expose only the correct data and concurrency primitives for
// 	  multiple concurrent processes to use.

// When working with Adhoc confinement, it becomes hard fairly quickly. As more engineers touch the code, the harder it
// gets to follow a convention that is followed within the codebase.
// On the other hand, lexical confinement makes it impossible to do the wrong thing, as it forces a channel within a
// scope to be initialized, written to and closed.

func main() {
	// The best approach to lexical confinement is through the channel ownership pattern.

	// Instantiates lexical scope for where the channel will be initialized, written to and closed.
	// It confines the written aspect of this channel to prevent other goroutines from writing to it.
	numCreatorOwner := func() <-chan int {
		nums := make(chan int)
		go func() {
			defer close(nums)
			for i := 0; i <= 100; i++ {
				nums <- i
			}
		}()
		return nums
	}

	// Declare consumer which receives a read-only channel.
	// By declaring that the only usage we require is read access, the usage of the channel within is confined to the
	// function to only read access.
	numCreatorConsumer := func(nums <-chan int) {
		for n := range nums {
			fmt.Println(n)
		}
	}

	// Receive the read aspect of the channel, which consumer will use.
	// This confines the main goroutine to a read-only view of the channels.
	nums := numCreatorOwner()
	numCreatorConsumer(nums)

	// So what's the point? Why pursue confinement if we have the synchronization available to us?
	// More performance and also reduced cognitive load.
	// Synchronization comes with a cost, if you can avoid it while preserving integrity and correctness, then do it.
}
