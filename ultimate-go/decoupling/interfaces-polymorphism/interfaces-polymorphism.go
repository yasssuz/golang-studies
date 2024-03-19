package main

import "fmt"

// Interfaces are not used to define concrete data types, but to define an abstract set of behaviors.
// Subsequently, those behaviors can be used by concrete data types (like a struct).
// Important to note that interfaces are value-less. They do not contain data that can be read or changed.
type Reader interface {
	// When creating an API design interface, make sure you're following those 2 important things:
	// - Your semantics (value and pointer) need to be consistent
	// - Your code needs to be sympathetic - both end-user and garbage collector
	read(b []byte) (int, error)
}

type File struct {
	name string
}

type Pipe struct {
	name string
}

func (file File) read(b []byte) (int, error) {
	s := "<rss><channel><title>Golang!</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

func (pipe Pipe) read(b []byte) (int, error) {
	s := `{name: "bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

// Retriever is accepting R based on what it can do, and not what it is.
// This is a very important part of interfaces, function can work based on
// behavior and not based on what value they have.
func retriever(r Reader) error {
	data := make([]byte, 100)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

func main() {

}
