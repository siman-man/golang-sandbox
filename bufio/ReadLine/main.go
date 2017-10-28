package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func readLine() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	line := readLine()
	fmt.Println(line)
}
