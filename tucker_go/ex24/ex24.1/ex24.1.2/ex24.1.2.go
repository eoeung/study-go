package main

import (
	"fmt"
	"sync"
)

// sync 패키지의 WaitGroup 설명
func exWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(3) // 작업 개수 설정
	wg.Done() // 작업이 완료될 때마다 호출
	wg.Wait() // 모든 작업이 완료될 때까지 대기
}

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done() // 작업이 완료됨을 표시 → wg의 남은 작업 개수를 1씩 줄임
}

func main() {
	// [서브 고루틴이 종료될 때 까지 기다리기]
	// - 함수 시간을 계산해서 time.Sleep()으로 대기하는 방법도 있지만, 함수 실행 시간을 늘 알 수 없음
	//   → sync 패키지의 WaitGroup 사용

	var goroutinCount = 30
	wg.Add(goroutinCount) // 총 작업 개수 설정 → 고루틴 생성 개수
	for i := 0; i < goroutinCount; i++ {
		go SumAtoB(1, 1000000000) // 각 루틴에서 SumAtoB()를 실행 → 30개의 고루틴으로 실행
	}

	wg.Wait() // 모든 작업이 완료되길 기다림 → 남은 작업 개수가 0이 되면 Wait()가 끝나면서 다음 줄로 넘어감
	fmt.Println("모든 계산이 완료됐습니다.")
}
