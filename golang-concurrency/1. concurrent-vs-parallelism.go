package main

func main() {
	// The difference between concurrency and parallelism is simple.

	// Concurrency -> a property of the code
	// Parallelism -> a property of the runtime

	// Concurrency means that a code runs out-of-order, and not necessarily in parallel.
	// Concurrent code is just code that will fork out of the main branch of execution and eventually run, but we do
	// not know when it will execute - this depends on the runtime. We program concurrent code hoping that it will run
	// in parallel.

	// Parallelism means branches of execution running at the same time. Maybe on different hardware threads, maybe on
	// inside a single hardware thread within an os thread period.
}
