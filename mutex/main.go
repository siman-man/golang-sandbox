package main

import (
	"github.com/siman-man/golang-sandbox/mutex/bank"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bank.Deposit(100)
			b := bank.Balance()
			fmt.Printf("%d\n", b)
		}()
	}

	wg.Wait()
}
