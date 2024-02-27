package main

import "fmt"

// In Go, the compiler performs a process called "escape analysis" to determine whether a
// variable's lifetime extends beyond the scope where it was created on the stack.
// If it does, the variable will "escape" to the heap instead.

type User struct {
	name  string
	email string
}

func main() {
	user1 := createUserV1() // Gets copy of value
	user2 := createUserV2() // Gets copy of heap address

	// We should always try to keep the data on the stack, as the stack is
	// self cleaning, more performant (specially around multi-threaded software), while
	// the stack has manual memory management and is less performant. Usually more latency
	// becasue the garbage collector gets involved

	fmt.Println("User 1:", user1)
	fmt.Println("User 2:", user2)
}

// Here, user is created as a value (not a pointer) inside the function createUserV1().
// When this function returns, the user variable will be copied to the caller's stack frame.
// Since it's a value and not a pointer, there's no need for heap allocation, and it won't escape.
func createUserV1() User {
	user := User{
		name:  "Bill 1",
		email: "bill1@gmail.com",
	}

	return user
}

// In this version, user is created as a local variable inside createUserV2(), but instead of returning the value itself,
// it returns a pointer to the local variable's address. Here, user would escape to the heap because it's being
// accessed outside the function's scope. Even though the function returns a pointer to the local variable, it's not safe to
// use this pointer outside the function because the variable user is allocated on the stack
// and will be deallocated once the function returns. This would result in a dangling pointer, causing undefined behavior.
func createUserV2() *User {
	user := User{
		name:  "Bill 2",
		email: "bill2@gmail.com",
	}

	// That's an allocation, when a value is constructed on the heap
	return &user
}

// This is an exapmle of the bad function above.
func createUserV2WrongWay() *User {
	// Now, we are using pointer semantic construction, which is bad.
	// This constructor basically says, create this struct and give me it's address.
	// By doing this, we completely lost visibilty on the `return`
	user := &User{
		name:  "Bill 2",
		email: "bill2@gmail.com",
	}

	// Is bad because, by looking into the line above only, it looks like we are using
	// value semantics, which would just return a copy to the above scope in the stack.
	// However, as we created user using pointer semantic construction, user is not being
	// copied, instead, is being escaped to the heap. This is hiding a latency cost.
	return user
}
