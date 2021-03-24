package main

import (
	"fmt"
	"sort"
)

func main() {
	pts := []Point{ // using Points will work, but we want to use the most basic type in general
		{10, 3},
		{3, 7},
		{6, 1},
	}

	//sort.Sort(pts) // won't work
	sort.Sort(Points(pts)) // type conversion
	fmt.Println("pts by X: ", pts)

	sort.Slice(pts, func(i, j int) bool { return pts[i].Y < pts[j].Y })
	fmt.Println("pts by Y: ", pts)
}

// Can't add methods to basic types
// func (p []Point) Len() int { return len(p) }

type Points []Point

func (p Points) Len() int { return len(p) }
func (p Points) Less(i, j int) bool { return p[i].X < p[j].X} // sort by X
func (p Points) Swap(i, j int) { p[i], p[j] = p[j], p[i]}

type Point struct {
	X int
	Y int
}

// minimal interface for sorting
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
