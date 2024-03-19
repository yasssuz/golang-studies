package main

type MyInterface interface{}

func main() {

	// IMPORTANT ABSOLUTES:
	// 1. NEVER COPY A VALUE COMING FROM A POINTER.
	//    When working with pointers, there is only one copy of the value, we don't want any other copies.
	//    From value-to-pointetr, in small situations, ok. Pointer-to-value; NEVER EVER EVER.

	// *******************************************************************************************

	// built-in types
	var i int
	var f float64
	var b bool
	var str string

	// With built-in types, we only move data around our program using value semantics.
	// Which means, we don't want to see pointers or addresses to any built-in types.
	// This is because of two reasons:
	// - Go team prefers to have those values in an isolate manner, without direct mutation.
	// - Those types have been created to be effeciently copied around our application.

	// Exception: if you need a `nil` value

	// *******************************************************************************************

	// Reference types
	var slc []string
	var m map[string]int
	var ch chan int
	var fn func(int, int) int
	var itf MyInterface

	// With reference types, we always pass them around our application using value semantics.
	// However, reference types have an interesting duality. You pass them around your application
	// using value semantics, but when you're reading or writing, under the hood is actually using
	// pointer semantics. This is beacuse, they refer to storage location rather than containing their own
	// data (for heavy things), and contain only very cheap information which is cheap to copy.

	// Exception: when using decode or marshall, you can pass a pointer.

	// *******************************************************************************************

}
