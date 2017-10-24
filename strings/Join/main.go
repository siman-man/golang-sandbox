package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{ "hello", "world" }
	fmt.Println(strings.Join(data, ","))
}
