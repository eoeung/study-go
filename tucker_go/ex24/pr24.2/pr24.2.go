package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup

type Account struct {
	Balance int
}

func DepositAndWithDraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance Should not be negative value: %d", account.Balance))
	}

	account.Balance += 1000
	account.Balance -= 1000
	fmt.Println(account.Balance)
	wg.Done()
}

func main() {
	account := &Account{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			DepositAndWithDraw(account)
		}()
	}
	wg.Wait()
}
