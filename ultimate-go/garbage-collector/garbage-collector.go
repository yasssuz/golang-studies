package main

func main() {
	// Go's garbage collector is:
	// - Non generational -> it doesn't divide the heap into separate generations
	// - Non compacting -> memory on the heap is not moving, its fixed for the length of the lifetime of the data stored
	// - Concurrent -> part of the GC is running at the same time with other applications
	// - Try color -> based on the tri-color marking algorithm
	// - Mark & Sweap -> based on mark-and-sweep algorithm to reclaim unreachable memory

	// Each garbage collector has different focus:
	// - Performance: can start and finish the collection as fast as possible
	// - Throughput: balancing the time spending on GC
	// - Resource: balancing the amount of the same memory resource used
	// etc...

	// Go's garbage collector focus on Throughput and Resources.
	// Go's was developed knowing that it most of the process was going to run on the cloud, and because of that
	// they focused a lot into using as least resources as possible to run the processes.

	// When the variable GOGC = 100, which is when the heap consumes the first 4meg of memory, the GC gets involved.
	// GC is divided into 3 steps:
	// - Mark start -> . Blocker.
	// - Marking -> Understanding what values are in use and those which are not. Concurrent.
	// - Mark termination -> . Blocker.
	// By putting all those 3 phases together, we are able to calculate how much time GC took.
	// Usually, is between 200 micro-seconds and 1 milli-second

	// The first and the last phase are blockers, which means, no process is running at that time.
	// GOGC tries to keep both of these processes under or at 100 micro-seconds.
	// On the other hand, Marking process is concurrent and will be running parallel to other processes.
	// However, is hard to predict how much time this will take as it depends on how much memory is being
	// used and what we are doing.

	// After "Mark start", when the processes come back to being executed, GOGC will need some CPU to continue to run it's processes.
	// GOGC will take exactly 25% of CPU available to do its "Marking" job. So, in a 4 core CPU, with 4 go routines running, while
	// on "Marking" phase, one of the cores will be assigned to garbage collection, and the other 3 routines will contiune with
	// application processes.
}
