package main

import (
	"fmt"
	"log"
	"unsafe"
)

const (
	MaxX = 1000
	MaxY = 1000
)

func main() {
	var p Player
	fmt.Println("player size:", unsafe.Sizeof(p))

	/* Before Object
	p = Player{"Parzival", 10, 20}
	fmt.Println("p =", p)
	fmt.Printf("%%+v: %+v\n", p)
	fmt.Printf("%%#v: %#v\n", p) // used in logging
	*/

	p1, err := NewPlayer("Parzival", 100, 200)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("p1", p1)

	p1.Move(700, 873)
	fmt.Println("p1 (move)", p1)

	if err := GatherAll([]mover{&p1, &Player{}, &Car{}}, 35, 3); err != nil {
		log.Fatalf("can't move: %s", err)
	}

	/*
		var m1 mover = &p1
		fmt.Println(m1)
	*/

	p1.FoundKey(Jade)
	fmt.Println("keys:", p1.Keys)
	p1.FoundKey(Crystal)
	fmt.Println("keys:", p1.Keys)
}

type mover interface {
	Move(int, int) error
}

func GatherAll(pieces []mover, x, y int) error {
	// FIXME: transaction (see STM)
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

func (o *Object) Move(x, y int) error {
	// FIXME: Bounds check
	o.X = x
	o.Y = y
	return nil
}

type Car struct {
	Object
	Length int
}

func NewPlayer(name string, x, y int) (Player, error) {
	if name == "" {
		return Player{}, fmt.Errorf("empty name")
	}
	if x < 0 || x > MaxX || y < 0 || y > MaxY {
		return Player{}, fmt.Errorf("(%d, %d) out of bounds (%d, %d)", x, y, MaxX, MaxY)
	}

	/*
		var p Player
		p.Name = name
		p.X = x
		p.Y = y
	*/

	p := Player{
		Name:   name,
		Object: Object{x, y},
	}
	return p, nil
}

type Key byte // byte is an alias to uint8 (rune = int32, a Unicode code-point/character)

const (
	Jade    Key = 1 << iota // 1<<0 = 1 (0001)
	Copper                  // 1<<1 = 2 (0010)
	Crystal                 // 1<<2 = 4 (0100)
)

// implement the fmt.Stringer interface
func (k Key) String() (result string) {
	keys := []Key{Jade, Copper, Crystal}
	for _, key := range keys {
		r := k & key
		if r > 0 {
			if len(result) > 0 {
				result += "|"
			}
			result += getKeyString(r)
		}
	}
	return result
}

func getKeyString(k Key) string {
	switch k {
	case 0:
		return "<nil>"
	case Jade:
		return "Jade"
	case Copper:
		return "Copper"
	case Crystal:
		return "Crystal"
	}

	// Careful: don't use %s or %v verbs here
	return fmt.Sprintf("unknown key: %d", k)
}

func (p *Player) FoundKey(k Key) {
	p.Keys |= k
}

type Player struct {
	Name string
	//	X      float64 // will "shadow" the Object.X
	Object // struct embedding
	Keys   Key
}

/*
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Reader interface {
    Read(n int) (p []byte, err error)
}
*/
