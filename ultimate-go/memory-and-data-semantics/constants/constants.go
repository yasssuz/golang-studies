package main

func main() {
	// Constants are not variables, they only exist at compile time
	// Constanst can be either of a type or of a kind

	// Untyped constants, minimun precision is 256 bits, which makes them mathematically precise
	const k1 = 500     // kind: integer
	const k2 = 73.3249 // kind: floating-point

	// Typed constants, precision/size comes from type itself
	const t1 int = 500         // type: int
	const t2 float64 = 73.3249 // type: float64

	// Variable of type float64
	var x1 = 1 / 0.333 // KindFloat(1) / KindFloat(0.333)
	// Note: a float can never be converted into a int, however, an int can be converted into a float

	// Constant of kind floating-point
	const x2 = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

	// Costant of kind integer
	const x3 = 1 / 3 // KindInt(1) / KindInt(3)

	// Represents a group of consts
	const (
		// Max integer value, will compile as const are 256 bits and not 64
		maxInt = 9223372036854775807

		// Way above 64 bits, will compile as consts are 256 bits
		biggerInt = 230368547758079223372036854775

		// Way above 64 bits, will NOT compile as is overflowing max amount of bits for an integer
		biggerButLimited int64 = 230368547758079223372036854775
	)
}
