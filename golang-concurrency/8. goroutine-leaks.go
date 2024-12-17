package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Goroutines are cheap and easy to create, and that's one of the main things that make Go's productivity so good.
// However, even tho goroutines are cheap, they do have a cost - starting at 4kb per goroutine - and they are not GC
// collected at runtime. So, regardless of how small their memory footprint is, we don't want to leave them lying about
// our process.
// Tip: Whoever creates the goroutine, should be able responsible for closing the goroutine.

func main() {
	//goroutineLeak()
	//goroutineLeakPreventionThroughDone1()
	goroutineLeakPreventionThroughDone2()
}

func goroutineLeak() {
	// In this example, the main goroutine passes a nil channel to doWork - causing the for range to never finish.
	// The goroutine will eventually get into an "awaiting" state, stop consuming any CPU but still leaving a memory
	// footprint.
	// In this example, the lifetime of the process is very short, but in a real program, goroutines could easily be
	// started at the beginning of a long-live program. In the worst case, the main goroutine could continue to spin up
	// goroutines through its life, causing creep in memory utilization.

	doWork := func(strings <-chan string) chan any {
		finished := make(chan any)
		go func() {
			defer close(finished)
			defer fmt.Println("goroutine ended")
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return finished
	}

	doWork(nil)
	fmt.Println("work done")
}

func goroutineLeakPreventionThroughDone1() {
	// In this example, we fix the goroutine leak by using the "done" convention.
	// We pass the done channel to the doWork function - as a first parameter, per convention.
	// We use a for-select pattern to either print, or exit the for-loop (if done has been signaled).
	// Another goroutine is created and will signal (close) the "done" channel after 5 seconds.

	doWork := func(done <-chan any, strings <-chan string) chan any {
		finished := make(chan any)
		go func() {
			defer close(finished)
			defer fmt.Println("goroutine ended")
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return finished
	}

	done := make(chan any)
	doWorkFinished := doWork(done, nil)

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Cancelling doWork goroutine")
		close(done)
	}()

	<-doWorkFinished
	fmt.Println("work done")
}

func goroutineLeakPreventionThroughDone2() {
	// In this example, the streamOwner function gets blocked while writing instead of while reading. -
	// We used the same approach of before - the "done" pattern, to tell the streamOwner when is time to close the
	// goroutine.

	randIntStreamOwner := func(done <-chan any) chan int {
		ints := make(chan int)
		go func() {
			defer close(ints)
			defer fmt.Println("goroutine closed")
			for {
				select {
				case ints <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return ints
	}

	done := make(chan any)
	randIntStreamConsumer := func(ints <-chan int) {
		for i := 1; i <= 3; i++ {
			fmt.Println(i, " ", <-ints)
		}
		close(done)
	}

	ints := randIntStreamOwner(done)
	randIntStreamConsumer(ints)

	time.Sleep(time.Second)
	fmt.Println("Work concluded")
}
