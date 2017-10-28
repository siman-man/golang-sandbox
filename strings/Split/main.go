package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	str := "-100, 1, 100"
	num_list := []int{}
	nums := strings.Split(str, ",")
	fmt.Println("Length = %d", len(nums))

	for index, n := range nums {
		if n, err := strconv.Atoi(strings.TrimSpace(n)); err == nil {
			num_list = append(num_list, n)
			fmt.Printf("%d: %d\n", index, n)
		}
	}
}
