package main

import "fmt"

func main() {
	c := 10
	list := make([]int, 0, c)

	for i := 0; i < 100; i++ {
		list = append(list, i)

		if c < cap(list) {
			c = cap(list)
			fmt.Printf("len = %d, cap = %d\n", len(list), cap(list))
		}
	}
}
