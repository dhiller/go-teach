package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p Player
	fmt.Println("player size:", unsafe.Sizeof(p))

	p, err := NewPlayer("Parzival", 10, 20)
	fmt.Println("p = ", p)
	fmt.Println("err = ", err)
	fmt.Printf("%%+v: %+v\n", p)
	fmt.Printf("%%#v: %#v\n", p) // used in logging
	fmt.Println("player size:", unsafe.Sizeof(p))

	p.Move(200,300)
	fmt.Println("p = ", p)

	c := Car{10, Object{17, 42}}
	fmt.Println("c = ", c)
	c.Object.Y = 37
	c.X = 27
	fmt.Println("c = ", c)

	player := Player{"test2", Object{99, 88}}
	car := Car{22, Object{77, 66}}
	fmt.Println("GatherAll = ", GatherAll([]mover{&p, &Player{}, &player, &car}, 22, 33))
}

type mover interface {
	Move(int, int) error
}

func GatherAll(pieces []mover, x, y int) error {
	// FIXME: transaction
	for _, p := range pieces {
		if err := p.Move(x, y); err != nil {
			return err
		}
	}
	return nil
}

type Object struct {
	X int
	Y int
}

type Car struct {
	Length int
	// X int // you could shadow the X in the struct
	Object // embedding
}

func (p *Object) Move(x,y int) error {
	p.X = x
	p.Y = y
	return nil
}

func NewPlayer(name string, x, y int) (Player, error) {
	return Player{name, Object{x, y}}, nil
}

type Player struct {
	Name string
	Object // embedding
}
