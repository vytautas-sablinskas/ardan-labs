package main

import (
	"fmt"
	"slices"
)

const (
	maxX = 600
	maxY = 400
)

type Item struct {
	X int64
	Y int64
}

type Key byte

const (
	Copper Key = iota + 1
	Jade
	Crystal
)

type Player struct {
	Name string
	Keys []Key
	X    int64
	Y    int64
	Item // embeded struct
}

func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	default:
		return fmt.Sprintf("unknown key - %d", k)
	}
}

func main() {
	a, b := 1, "1"
	fmt.Printf("a=%v, b=%v\n", a, b)
	fmt.Printf("a=%#v, b=%#v\n", a, b)

	var i Item
	fmt.Printf("i: %#v\n", i)

	i = Item{10, 20} // must specify all fields in order
	fmt.Printf("i: %#v\n", i)

	// can be in any order, can be partial
	i = Item{
		X: 15,
		//Y: 25,
	}
	fmt.Printf("i: %#v\n", i)

	// factory method to ensure necessary values are injected / correct
	i, err := NewItem(30, 25)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("i: %#v\n", i)

	i.Move(10, 10)
	fmt.Printf("i: %#v\n", i)

	p1 := Player{
		Name: "test",
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	p1.Move(100, 200)
	fmt.Printf("p1: %#v\n", p1)

	err = p1.Found(Jade)
	if err != nil {
		fmt.Println("err:", err)
	}
	err = p1.Found(Key(0))
	if err != nil {
		fmt.Println("err:", err)
	}
	err = p1.Found(Jade)
	if err != nil {
		fmt.Println("err:", err)
	}
	err = p1.Found(Copper)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(p1.Keys)

	moveAll([]Mover{&p1, &i}, 10, 20)
	fmt.Printf("p1.X - %d, p1.Y - %d\n", p1.X, p1.Y)
	fmt.Printf("i.X - %d, i.Y - %d\n", i.X, i.Y)
}

func NewItem(x, y int64) (Item, error) {
	if x < 0 || x > maxY {
		return Item{}, fmt.Errorf("X is out of bounds - %d", x)
	}

	if y < 0 || y > maxY {
		return Item{}, fmt.Errorf("Y is out of bounds - %d", y)
	}

	return Item{
		X: x,
		Y: y,
	}, nil
}

// i is the receiver
// i is copied, so can't modify it directly, unless the receiver is pointer to item
func (i *Item) Move(dx, dy int64) {
	i.X += dx
	i.Y += dy
}

// Exercise, error if key is not one of "jade", "copper", "crystal". Should add key only once
func (p *Player) Found(key Key) error {
	validKeys := []Key{Copper, Crystal, Jade}
	if !slices.Contains(validKeys, key) {
		return fmt.Errorf("Invalid key: %s", key.String())
	}

	if slices.Contains(p.Keys, key) {
		return fmt.Errorf("Key was already found: %s", key.String())
	}

	p.Keys = append(p.Keys, key)

	return nil
}

func (p *Player) Move(dx, dy int64) {
	p.X += dx
	p.Y += dy
}

type Mover interface {
	Move(dx, dy int64)
}

func moveAll(movers []Mover, dx, dy int64) {
	for _, m := range movers {
		m.Move(dx, dy)
	}
}
