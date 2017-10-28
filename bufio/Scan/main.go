package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Printf("%q\n", scanner.Text())
	}
}
