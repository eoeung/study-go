package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock() // 뮤텍스 획득 → 다른 고루틴이 뮤텍스를 획득한 경우, 뮤텍스를 놓을 때 까지 대기

	// mutex.Lock()을 통해 뮤텍스를 획득한 하나의 고루틴만 아래 로직을 실행
	defer mutex.Unlock() // defer를 사용한 Unlock()
	// 한 번 획득한 뮤텍스는 반드시 Unlock()을 호출해서 반납해야함

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance Should not be negative value: %d\n", account.Balance))
	}

	time.Sleep(time.Microsecond)
	account.Balance += 1000

	time.Sleep(time.Microsecond)
	account.Balance -= 1000

	fmt.Printf("### %v\n", account)
}

func main() {
	// [뮤텍스를 이용한 동시성 문제 해결]
	// - 단순하게 한 고루틴에서 값을 변경할 때, 다른 고루틴이 건들지 못하게 하는 것
	// - 뮤텍스(mutex)를 이용해서 자원 접근 권한을 통제할 수 있다.
	// - mutual exclusion의 약자 → 상호 배제
	// - 자원 접근 권한이라는 역할을 수행
	// - Lock()메서드로 뮤텍스 획득 / UnLock()메서드로 뮤텍스 반납

	var wg sync.WaitGroup

	account := &Account{}
	var goroutineCount = 10
	wg.Add(goroutineCount)
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
