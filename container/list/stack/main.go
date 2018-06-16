package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := list.New()
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	for stack.Len() > 0 {
		v := stack.Remove(stack.Back())
		fmt.Println(v)
	}
}
