package main

func main() {
	// It's common for languages to have their chain of abstraction at the level of the OS thread and
	// memory access synchronization. In those languages, you start taking into consideration technicalities that are
	// critical for your concurrency model to work, even before thinking at how to solve the problem itself.
	//Go takes a different route.

	// Goroutines (the way to handle concurrency in go) free us from having to think about our problem space in
	// terms of parallelism and instead allow us to model problems closer ot their natural level of concurrency.

	// Goroutines are extremely lightweight - so lightweight that thinking at how many you're going to have before
	// encountering a performance problem is considered premature optimization. On the other hand, in other languages
	// that work directly with threads, you need to consider those matters upfront.

	// CSP (or Communicating-Sequential-Progresses) is a technique and also the title of a research paper that is what
	// inspired go's concurrency primitives.

	// Those concurrency primitives (such as channels and select) are the recommended approach to work with concurrency.
	// However, Go also supplies the more traditional methods of writing concurrent code through memory access synchronization.

	// How to choose:
	// - Channels:
	// 		- If you have code that transfer ownership of data -> similar to memory-ownership concept of languages that
	//		  do not have a garbage collector: data has an owner, and we need to ensure only one concurrent context has
	//	      ownership of the data at the time.
	//      - Coordinating multiple pieces of logic -> channels are way more composable, and using select statement + their
	// 		  ability to serve as queues and be safely passed around. Having deadlocks and race-conditions may be the sign
	// 		  that you should be using channels.
	//
	// - Primordial way
	// 		- If you're trying to guard internal state of a struct -> for example, if you want to hide the implementation
	// 		  detail of locking your critical section from your callers.
	// 		- If is performant critical -> does not mean "I want my code to be faster", rather, if you have a section
	//		  of your code that you have profiled and seems to be causing significant performance issues, you should consider
	// 		  switching to the primordial way - although you actually may need to consider a complete re-architecture.

}
