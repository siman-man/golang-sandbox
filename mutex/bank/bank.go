package bank

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	return balance
}

func Withdraw(amount int) bool {
	deposit(-amount)
	if Balance() < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {
	balance += amount
}
