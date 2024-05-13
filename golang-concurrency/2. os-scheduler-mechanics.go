package main

func main() {
	// concurrency -> code that executes out of order
	// parallelism -> things executing at the same time

	// We program concurrent code hoping that it will run in parallel

	// a CPU can have many cores, which can have many hardware threads.
	// Each one of those HTs is allowing us to run instructions on it.
	// They are managed exclusively by the CPU hardware and are invisible to the operating system.
	// Each HT can execute instructions independently and concurrently with other HTs on the same core.

	// On the other hand, system threads (or kernel threads) are managed by the operating system.
	// Each os thread has his own stack, program counter, etc.
	// They are scheduled by the os scheduler, which allocates CPU time to each.
	// They can run concurrently on multiple hardware threads if the CPU supports parallel execution.

	// an OS thread can have three states:
	// 1. Running state -> is executing the processes he was assigned to on an HT
	// 2. Runnable state -> is waiting its turn to run on the HT
	// 3. Awaiting state -> is waiting for processes to be assigned

	// Sometimes, parallelism is simulated by running multiple OS threads within a single scheduler period.
	// For example, given a 100ms scheduler period, you could split it as 50ms TA and 50ms TB.
	// Actually, you can split the time inside the scheduler period as much as you want, even for 100 OS threads.
	// If you split 100 OS threads, 1ms each, within a given scheduler period of 100ms, you would have 1ms (or 1000ns)
	// per OS thread. This is around 12K instructions to be executed.

	// Now, changing between OS threads on the scheduler does not come for free. Every time we change, something called
	// "Context Switch" happens, and it costs anywhere between 1ms to 2ms (or 1000ns to 2000ns). This may not seem like
	// a lot, but its around 12K to 24K instructions! It's not feasible anymore to change between contexts, there needs
	// to be a minimum threshold where it is going to be worth it to change between contexts.

	// Another important subject is workload, and we have two of them:
	// CPU-bound -> workload were a thread naturally doesn't move into an awaiting state.
	// IO-bound -> aka blocking workload, is when the thread never actually uses its full time slice.
	// 		This is because, while doing a process, it may make a call to an API and the thread will go into awaiting state.

	// So, on CPU-bound workloads, those context switches are our enemy, as they keep hitting 12/24K or lost processes every time.
	// On IO-bound workloads, on the other hand, context switches allow us to keep certain threads into an awaiting state and
	// assign more threads, which allows us to get things done faster.

	// In Go, there are 4 type of events that can cause a context switch:
	// 	1. Creating a Goroutine
	//	2. Garbage collection
	//	3. Syscall (synchronous and asynchronous)
	//	4. Blocking operations (like synchronization)
}
