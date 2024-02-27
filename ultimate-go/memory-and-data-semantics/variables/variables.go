package main

func main() {
	// Variable: providing a name to some memory allocation to be able to reference it
	zeroValueVariables()
	valueVariables()
}

func zeroValueVariables() {
	// Declare variables (with built-in types) that are set to their "zero-value"
	// "zero-value" in Golang means that once memory is allocated, it needs to be initialized.
	// if it is not initialized (for instance, not allocating an initial value), the compiler will
	// occupy the memory slot with zeros.

	// Represents a number without decimals, "zero-value" is 0
	var a int
	// int8 -> 1 byte, hold values from -128 to 127
	// int16 -> 2 bytes, hold values from -32768 to 32767
	// int32 -> 4 bytes, hold values from -2147483648 to 2147483647.
	// int64 -> 8 bytes, hold values from -9223372036854775808 to 9223372036854775807
	// int -> type's size depends on the underlying hardware. On a 32-bit hardware, it is int32, and on a 64-bit hardware, int64.

	// Represents a string, "zero-value" is " "
	var b string
	// Go has an unique implementation for strings, it has a two "word" value:
	// - Pointer, "zero-value" is nil
	// - Int, "zero-value" is 0

	// Represents two things:
	// - IEEE 754 Binary decimal
	// - It occupies 64 bits of space (8 bytes)
	// "zero-value" is 0
	var c float64
	// float in golang is uses IEEE 754 standard
	// float32 -> single-precision floating-point numbers, 7 decimal digits
	// float64 -> fdouble-precision floating-point numbers, 15 decimal digits
	// "floating-point" number = numbers that can have fractional part

	// Represents a boolean, which is 1 bit (and a full byte) of allocation
	// "zero-value" is false
	var d bool
}

func valueVariables() {
	// Declare variables and initialize them, at the same time
	// Type is assigned automatically

	// Represents type int
	aa := 10

	// Represents type string
	bb := "hello"
	// Two word string value is:
	// [h] [e] [l] [l] [o]
	// - Pointer pointing at [h]
	// - Int with value 5

	// Represents type float64
	cc := 3.343

	// Represents type bool
	dd := true
}
