package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	// [동시성 프로그래밍 주의점]
	// - 고루틴은 별도의 CPU 코어에서 동작하지만, 같은 메모리 공간에 동시에 접근이 가능하다.

	// Ex) 같은 자원에 여러 고루틴이 접근해서 동시성 문제를 일으키는 예
	var wg sync.WaitGroup
	var goroutineCount = 10
	wg.Add(goroutineCount)

	account := &Account{}

	for i := 0; i < goroutineCount; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 { // 잔고가 0 미만이면 패닉
		panic(fmt.Sprintf("Balance Should not be negative value: %d\n", account.Balance))
	}

	time.Sleep(time.Millisecond)
	account.Balance += 1000 // 1000원 입금

	time.Sleep(time.Millisecond)
	account.Balance -= 1000 // 1000원 출금

	fmt.Printf("### %v\n", account)
}

// 1) 순차적으로 0 → +1000 → 1000 - 1000 → 0이 되어야 맞는데, 각각의 고루틴이 같은 메모리에 동시에 접근
