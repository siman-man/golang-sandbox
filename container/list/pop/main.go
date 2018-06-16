package main

import (
	"fmt"
	"container/list"
)

func main() {
	type Coord struct {
		Y int
		X int
	}
	que := list.New()
	que.PushBack(Coord{Y: 1, X: 2})
	que.PushBack(Coord{Y: 0, X: 1})
	que.PushBack(Coord{Y: 3, X: 3})

	for que.Len() > 0 {
		coord, ok := que.Remove(que.Front()).(Coord); if ok {
			fmt.Printf("y = %d, x = %d\n", coord.Y, coord.X)
		}
	}
}
