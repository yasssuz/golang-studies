package main

import "fmt"

type Data struct {
	name string
	age  int
}

// In Golang, is a bad practice to create "getters" and "setters" methods as in common
// OOP languages. Getters and setters are NOT an API in go, APIs must provide something, and
// just changing state is not enough. The examples below are just illustrative.

func (d Data) displayName() {
	fmt.Printf("My name is %#v", d.name)
}

func (d *Data) setAge(age int) {
	d.age = age
	fmt.Printf("My new age is %d", age)
}

func main() {

	d := Data{name: "Karim"}

	fmt.Println("Proper calls to method")

	// How we call methods in Go
	d.displayName()
	d.setAge(19)

	// ***************************************************************************************8
	f1 := d.displayName

	d.name = "Madalina"

	f1() // Will still display "Karim"

	// This is because we are working with value semantics, and at the time we
	// created that function (without calling) we made our own copy of d and stored it locally.
	// `f1` has is a two words data structure:
	// - pointer -> pointing to the code
	// - pointer -> to the copy of `d`
	// This can cause bugs, and it also causes an allocation on the heap.

	f2 := d.setAge

	d.name = "Buni"

	f2(32)

	// In function `f2`, the new changed name would be displayed. This is beacuse `setAge`
	// is using pointer semantics.
	// `f2` looks like:
	// - pointer -> pointing to the code
	// - pointter -> pointing to the original `d` struct, and not pointing to its own copy.
	// In this case, we shouldn't have bugs. Unfortunately tho, because of GO's escape analysis
	// not being perfect, this will also cause an allocation on the heap.

	// ***************************************************************************************8

	// So, when working with methods in go, is important to remember that they will cause allocation
	// on the heap and indirection, regardless of the data semmantics we are working with.
	// If you want to do the fastest possible Go code, you can't be decoupling the code. Allocations on
	// the heap will cause garbagge collector to be involved, which will create latency.

	// In the end of the day, we are not trying nto create the fastest program as possible, but a program
	// which is just fast enough. Creating productive allocation which is going to make our life easier is the correct
	// path to follow.
}
