package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	// [자원 관리 기법]
	// - 뮤텍스가 좋은 방법이지만 뮤텍스만 사용하게 된다면 동시성 프로그래밍을 사용하지 않는 것과 같음
	// - 모든 문제는, 같은 자원을 여러 고루틴이 접근하기 때문에 발생

	// 1) 영역을 나눈다.
	// 2) 역할을 나눈다.

	// Ex) 영역을 나누는 방법

	var jobList [10]Job // Job 타입 배열 선언, 길이는 10

	for i := 0; i < len(jobList); i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(len(jobList))

	for i := 0; i < len(jobList); i++ {
		fmt.Printf("### i :: %d\n", i)
		job := jobList[i]
		go func() { // 각 작업을 고루틴으로 실행
			job.Do()
			wg.Done()
		}()
		// → 각 고루틴은 할당된 작업만 하므로, 간섭이 발생하지 않음
		//   Ex) 파일 100개를 읽을 때는 파일별로 고루틴을 할당해서 수행
		// ∴ 고루틴간 간섭을 없애는 것이 가장 중요
	}
	wg.Wait()
}

/*
	### i :: 0
	### i :: 1
	### i :: 2
	0 작업 시작
	1 작업 시작
	### i :: 3
	2 작업 시작
	### i :: 4
	### i :: 5
	3 작업 시작
	### i :: 6
	### i :: 7
	### i :: 8
	### i :: 9
	7 작업 시작
	9 작업 시작
	5 작업 시작
	8 작업 시작
	6 작업 시작
	4 작업 시작
	4 작업 완료 - 결과: 16
	9 작업 완료 - 결과: 81
	0 작업 완료 - 결과: 0
	3 작업 완료 - 결과: 9
	7 작업 완료 - 결과: 49
	8 작업 완료 - 결과: 64
	5 작업 완료 - 결과: 25
	6 작업 완료 - 결과: 36
	1 작업 완료 - 결과: 1
	2 작업 완료 - 결과: 4
*/
