package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("foobar", "bar"))
	fmt.Println(strings.Contains("foobar", "test"))
	fmt.Println(strings.Contains("seafood", ""))
}
