package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	ch <- "A"

	close(ch)

	for n := range ch {
		fmt.Printf("%s\n", n)
	}
}
