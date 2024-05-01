package main

func main() {
	// Packaging directly conflicts with how we have been taught to organize our code in other languages.
	// In other languages, packaging is a feature - something you can use or ignore.
	// In other languages, we usually use directories to define a mental model of where we save things.

	// In go, packaging is the basic structure of your project - everything should be seen as a package.
	// Packaging in go is like applying the idea of a microservice at the source tree level.
	// All packages are "first class", and the only hierarchy is what you define in the source tree of the package.
	// There needs to be a way to "open" parts of the package to the outside world.
	// Two packages cannot be cross-imported.

	// Packages in go should:
	// - Be purposeful -> packages should provide, not contain.
	// 		- If you're unable to explain what it does, then it has no purpose.
	// 		- Every package should be their island - packages like "models" or "utils" only contain.
	// - Be usable -> designed like a user is going to use it.
	// 		- Should be intuitive and clear - they should have docs.
	// 		- Should be safe - protecting user's app from cascading effects.
	// 		- Must make things simpler, not harder.
	// 		- To be safe - they must have tests.
	// - Be portable -> designed with reusability in mind
	// 		- Must use few dependencies.
	// 		- Should reduce setting policy when its practical.
}
