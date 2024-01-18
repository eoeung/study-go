package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// [뮤텍스와 데드락]

	// [뮤텍스의 단점]
	// 1) 뮤텍스로 동시성 프로그래밍의 문제를 해결했지만, 동시에 성능 향상이 안됨
	//   - 동시성 프로그래밍으로 성능 향상을 하려고 했는데 못하게 되는 아이러니가 발생

	// 2) 데드락이 발생할 수 있음
	//   - 데드락 : 어떤 고루틴도 원하는 만큼 뮤텍스를 확보하지 못해서 무한히 대기하게 되는 경우
	//   - 데드락은 프로그램을 완전히 멈추게 만든다.

	// Ex)
	//   1) 테이블 위에 숟가락과 포크가 각각 1개씩 존재
	//   2) A는 숟가락을, B는 포크를 집음
	//   3) 이제 A는 포크를, B는 숟가락을 집으려고 했지만 테이블에는 아무 것도 없음
	//   4) 서로 양보하지 않아 A, B 둘 다 밥을 먹을 수 없음

	rand.Seed(time.Now().UnixMicro())

	var goroutineCount = 2 // A와 B
	wg.Add(goroutineCount)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "숟가락") // A는 '포크 - 숟가락' 순서로
	go diningProblem("B", spoon, fork, "숟가락", "포크") // B는 '숟가락 - 포크' 순서로
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)

		first.Lock() // 첫 번째 뮤텍스 획득 시도
		fmt.Printf("%s %s 획득\n", name, firstName)

		second.Lock() // 두 번째 뮤텍스 획득 시도
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// 뮤텍스 반납
		second.Unlock()
		first.Unlock()
	}
}
