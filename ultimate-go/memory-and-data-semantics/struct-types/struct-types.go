package main

// Struct types are the second type of types available in Go.
// A Struct type is the idea of defining the type you want, modeling the data yourself.

type Example1 struct {
	flag    bool
	counter int16
	pi      float32
}

// Memory allocation: [ [ 1 byte ] [ 2 bytes ] [ 4 bytes ] ] = total 7 bytes? NOPE.
// Is 8, because of something called "alignments".
// There needs to be a "padding-byte" between the bool (1 byte) and the int16 (2 bytes).

// If counter was int32 (4 bytes), then the padding-byte would be 3.
// Similarly, if counter was int64 (8 bytes),the the padding-byte would be 7.

// However, here's another problem: the largest field (int64 8bytes) needs to fall
// into an alignement with the entire struct (20 bytes). Well, 20 is not divisible by 8.
// This means that we need to add more padding-bytes into the end until it fits the alignment.
// 24 bytes.

// Here's the good news, you can reduce the padding by simply putting the biggest values at the start:
type Example2 struct {
	counter int64
	pi      float32
	flag    bool
}

// Now, only padding at the end needs to be added, to get from 13 bytes to 16 bytes (which fits the alignment).
// Note: do this only if there is a performance need for it; don't do premature optimization.

func main() {
	// Declare a variable with type struct set to "zero-value"
	var e1 Example1

	// Declare a variable of type example and initialize using a struct literal
	e2 := Example1{
		flag:    true,
		counter: 5,
		pi:      89.3443,
	}

	// Declare a variable with an anonymous struct set to "zero-value"
	var e3 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Declare and initialize a variable with an anonymous struct
	e4 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 92,
		pi:      53.820,
	}
}
