package main

import (
	"fmt"
	"sync"
)

// Channels are one of the synchronization primitives in Go derived from Hoare's CSP.
// While they can be used to synchronize access of the memory, they are best used to communicate information between
// goroutines. Channels serve as a conduct for a stream of information; values may be passed along the channel, and then
// read out downstream.

func main() {
	// Create a channel of type any. This channel can both have a value written to it or read.
	ch1 := make(chan string)

	// Below are unidirectional channels - that is, channels that only support sending or receiving information.
	// Simply adding <- operator on the left makes it only readable, adding it to the right makes it only sendable.
	// ch2 := make(<-chan any)
	// ch3 := make(chan<- any)

	// It's not very common to see unidirectional channels instantiated, but often is common to see them used as
	// function parameters and return types, which is very useful, as we'll see. This is because Go will implicitly
	// convert bidirectional channels to unidirectional channels when needed.

	// To use channels, we use again the <- operator. We declare <- on the right to pass data, and on the left to read.
	go func() {
		ch1 <- "Hello world"
	}()
	fmt.Println(<-ch1)

	// Noticed that the value was printed correctly? Even without any guarantee that a concurrent function will run
	// before the printing?
	// Well, this is because channels in Go are blocking. Any goroutine that attempts to write to a channel that is full
	// will wait until the channel has been emptied, and any goroutine that attempts to read from a channel that is
	// empty will wait until at least one item is placed on it.
	// In the above case, `fmt.Println` contains a pull from the channel `ch1` and it will sit there until a value is
	// placed on the channel. Likewise, the anonymous goroutine is attempting to place a string into the channel, and so
	// the goroutine will not exit until the writing is successful. Thus, the anonymous goroutine and the main goroutine
	// block deterministically.

	fmt.Println("-----------------------------------------------------------------------------------------------")
	//closingChannel()
	//okIndicator()
	//bufferedChannel()
	channelOwnership()
}

func closingChannel() {
	// In programs, is very useful to be able to indicate that no more values will be sent over a channel.
	// This helps downstream read processes to know when to move on, exit re-open communications, etc.
	// To close a channel, you can use the `close` keyword.

	ch1 := make(chan int)

	go func() {
		defer close(ch1) // Ensure chanel is closed before we exit the goroutine.
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
	}()

	for num := range ch1 { // Range over ch1 -> will automatically break the loop when the given chanel is closed.
		fmt.Printf("%v \n", num)
	}

	// You can also use close to unblock awaiting goroutines simultaneously.

	ch3 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			defer wg.Done()
			<-ch3 // blocks goroutine until closure of chanel
			fmt.Printf("%v has begun \n", i)
		}()
	}

	fmt.Println("Unblocking goroutines...")
	close(ch3) // closes chanel and simultaneously unblock goroutines.
	wg.Wait()
}

func okIndicator() {
	ch1 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	// The second return value - "ok" - is a way for a read operation to indicate whether the read off the channel
	// was a value generated by a writing elsewhere in the process, or a default value generated from a closed channel.
	num, ok := <-ch1
	fmt.Printf("(%v): %v \n", ok, num)
	// In the above case, ok is true, and num is 1.

	// If we close the channel, the 'ok' return value will be false and the num will be 0.
	// This is because:
	// 1. If the closed channel was open and not empty, reads will succeed until the channel is drained (ok == true, num != 0).
	//    and when emptied will produce its default value (ok == false, num == 0).
	// 2. If the closed channel open and was empty, reads will be (ok == false, num == 0).
	close(ch1)
	num, ok = <-ch1
	fmt.Printf("(%v): %v \n", ok, num)
}

func bufferedChannel() {
	// A buffered channel is a channel that has a buffer of a certain size. The buffer size is specified when the
	// channel is created. The buffer size is the number of items that can be stored in the channel before the
	// channel blocks when a writing is attempted. Even if no reads are performed, a goroutine can still perform n writes.
	// If the buffer size is 0, the channel is unbuffered and will block when a writing is attempted.
	ch1 := make(chan rune, 2)

	// As we mentioned earlier, writes to channel that is full blocks the goroutine until the channel is not full.
	// Also, read from a channel that is empty blocks the goroutine until the channel is not empty.
	// An unbuffered channel has a capacity of 0, and thus is blocked before any writing.
	// On the other hand, a buffered channel has a capacity of n, and thus will get full after n writings.

	// Buffered channels are an in-memory FIFO queue for concurrent processes to communicate over.
	// Assuming that there are no reads, the letter 'a' will be written to the channel, and then the letter 'b'.
	ch1 <- 'a'
	ch1 <- 'b'
	// At this moment, if we try to write to the channel, it will block until the channel is not full anymore.
}

func channelOwnership() {
	// Ownership, in the context of channels, is defined by a goroutine that instantiates, writes and closes a channel.
	// Unidirectional channel declarations are the tool that will allow us to distinguish between goroutines that own
	// channels and those that only use them.

	// Channel owners have a write-access view into the channel (chan or chan<-), and channel consumers have a
	// read-only view of the channel (<-chan).
	// Once the distinction between channel owners and non-owners is made, the results from the preceding table follow
	// naturally, and we can begin to assign responsibilities to goroutines that own channels and those that do not.

	// Channel owners:
	//   1. Instantiate channel.
	//   2. Perform wries, or pass ownership to another goroutine.
	//   3. Close channel.
	//   4. Encapsulate the previous three things in this list and expose them via a reader channel.
	// By assigning these responsibility to the channel owner, we can ensure that:
	//   - As we initialize the channel -> we remove the risk of deadlocking by writing to a nil channel.
	//   - As we initialize the channel -> we remove the risk of panicking by closing a nil channel.
	//   - As we write and close the channel -> we remove the risk of panicking by writing to a closed channel.
	//   - As we close the channel -> we remove the risk of panicking by closing a channel twice.
	//   - We wield the type check at complete time to prevent improper writes to our channel.

	// Channel consumers:
	//   1. Knowing when the channel is closed -> by examining the second return value of the read operation.
	//   2. Responsibly handling blocking for any reason -> depends on your algorithm; you may want to stop reading or
	//      wait for a timeout, or you may just be content to block for the lifetime of the process.

	// Let's put this knowledge into practice.

	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5) // Instantiate a buffered channel with a capacity of 5 (we'll produce 6 results).
		go func() {
			defer close(resultStream) // Close the channel when we're done.
			for i := 0; i < 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream // Return the channel -> return value will be converted to read-only, as declared by chanOwner func.
	}

	consumerStream := chanOwner()

	for num := range consumerStream {
		fmt.Printf("Received: %v \n", num)
	}

	// The above code will print:
	// 0
	// 1
	// 2
	// 3
	// 4
	// This is because the channel is buffered, and the goroutine that created the channel will block until the channel
	// is not full anymore.
}