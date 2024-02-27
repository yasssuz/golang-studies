package main

func main() {
	// If a stack goes out of frames, the program doesn't just "shut down".
	// Actually, the size of the stack doubles. It starts at 2K, then 4k, then 8K, etc.

	// To double the size of the stack, you need to create a new stack with double the
	// size and move all the values from the previous stack to the new one.

	// Imagine, tho, that you had many stacks sharing values between each other, using pointers.
	// And then, the frames on a stack finish and you need to double it. Now, all the stacks that
	// are using a value from the stack that doubled in size, would all need to update and point to the new address.
	// *Latency screams...*
	// And this is why, values cannot be shared across stacks. What is created in the stack, stays in the stack.
	// The only way values can be shared across stacks is if the value is escaped to the heap, as the heap
	// doesn't need to "double itself". Everything on the heap is constant until explicitaly saying otherwise.

}
