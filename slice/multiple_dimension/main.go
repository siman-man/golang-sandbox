package main

import (
	"fmt"
)

func main() {
	s := make([][]int, 4)
	for i := 0; i < 4; i++ {
		s[i] = make([]int, 4)
	}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			fmt.Println(s[y][x])
		}
	}
}
