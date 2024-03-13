package main

import "fmt"

func main() {
	var i1 Item
	fmt.Printf("i1 -> %#v %d\n", i1, i1)

	i2 := Item{
		X: 1,
		Y: 2,
	}
	fmt.Printf("i2 -> %#v %d\n", i2, i2)

	i3 := Item{4, 5}
	fmt.Printf("i3 -> %#v %d\n", i3, i3)

	i4, err := NewItem(7, 80)
	fmt.Printf("i4 -> %#v %d\n error: %d\n", i4, i4, err)

	i4.Move(200, 400)
	fmt.Printf("i4 (move) -> %d\n", i4)

	p1 := Player{
		Name: "Karim",
		Item: Item{200, 300},
	}
	fmt.Printf("p1 -> %#v %d \n p1.X: %d", p1, p1, p1.X)
	p1.Move(400, 600)
	fmt.Printf("p1 (move) -> %d\n", p1)

	ms := []Mover{
		&i1,
		&i2,
		&i3,
		i4,
		&p1,
	}

	moveAll(ms, 0, 0)
}

func moveAll(ms []Mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

// i is called "the receiver", similar to "this"
// if you want to mutate, use pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func NewItem(x, y int) (*Item, error) {
	if x > maxX || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bonds %d/%d", x, y, maxX, maxY)
	}

	i := Item{x, y}

	// The Go-compiler will induce "escape analysis" and will allocate i to the heap
	return &i, nil
}

const (
	maxX = 1000
	maxY = 600
)

type Mover interface {
	Move(x, y int)
	// Move(int, int) -> without named parameters, also works
}

type Player struct {
	Name string
	Item // Embed item -> creates Item key based on Item struct, however, it also lifts up all Item keys to the
	// top level of Player. This means that, you can access Player.X instead of Player.Item.X.
	// An issue tho, is that if the struct has a conflicting key, the go compiler will not notify you of that and
	// this can create bugs.
	// Item Item -> you could also do this way, but you'd lose the features above mentioned.
}

type Item struct {
	X int
	Y int
}
