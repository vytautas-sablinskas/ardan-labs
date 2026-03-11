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

type Player struct {
	Name string
	Keys []string
	Item // embeded struct
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

	err = p1.Found("jade")
	if err != nil {
		fmt.Println("err:", err)
	}
	err = p1.Found("incorrect")
	if err != nil {
		fmt.Println("err:", err)
	}
	err = p1.Found("jade")
	if err != nil {
		fmt.Println("err:", err)
	}
	err = p1.Found("copper")
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(p1.Keys)
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
func (p *Player) Found(key string) error {
	validKeys := []string{"jade", "copper", "crystal"}
	if !slices.Contains(validKeys, key) {
		return fmt.Errorf("Invalid key: %s", key)
	}

	if slices.Contains(p.Keys, key) {
		return fmt.Errorf("Key was already found: %s", key)
	}

	p.Keys = append(p.Keys, key)

	return nil
}
