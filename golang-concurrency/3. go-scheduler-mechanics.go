package main

func main() {
	// Go's mechanism for handling goroutines is an implementation of an M:N scheduler, which means it maps
	// M green threads to N OS threads. Goroutines are then scheduled to green threads.

	// When there are more goroutines than green threads, the scheduler handles the distribution of the goroutines
	// across the available threads and ensures that when these goroutines become blocked, other can run.

	// Go follows a model of concurrency called the fork-join model. As you're going to read the next chapter,
	// a goroutine is simple a branch-of-execution, and to create this branch, we create a fork - or the point in the
	// program where a child branch of execution was split off and is now running concurrently to its parent.
	// The word "join" refers to the fact that at some point in the future, these concurrent branches
	// of execution will join back together.

	// Do you remember we talked about context switches at the os scheduler level?
	// At the go scheduler level, context switching is much cheaper, around 200 nanoseconds instead of the
	// standard 1000/2000ns.
	// This is because the runtime can be more selective in what is persisted for retrieval, how it is persisted and
	// when persisting needs to occur.

	// Also, while at the os level context switches are a somewhat unpredictable event, at the go level, they are
	// deterministic. This is because the scheduler has some safe-points to do those context switches,
	// and whenever a function call transition happens, the scheduler takes an opportunity to do a context switch.
	// Those safe-points are:
	// - `go` keyword -> when a goroutine is created.
	// - GC -> anytime the garbage collector runs.
	// - syscall -> like logging, file I/O, network I/O, etc.
	// - Certain blocking calls -> like synchronization, orchestration, mutexes, atomic instructions, CGO, etc.
}
