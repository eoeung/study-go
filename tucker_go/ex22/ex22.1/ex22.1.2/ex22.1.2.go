package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

// 요소를 반환하면서 삭제
func (q Queue) Pop() interface{} {
	// v.Front() // v.Back()
	// - 맨 앞 / 뒤 요소를 반환
	//   → 리스트가 비어있으면 nil 반환
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front) // 리스트 내에 요소를 삭제하고, 그 요소의 값을 반환
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{v: list.New()}
}

func main() {
	// [리스트로 큐 구현하기]
	// - 대기열 작업, 명령 큐 같이 순서가 유지되어야 하는 경우에 사용

	// - 큐의 특징
	//   1) 들어간 순서대로 빠져나오기 때문에, 순서가 유지됨
	//   2) 새로운 요소는 늘 마지막에 추가됨
	//   3) 출력값은 맨 앞에서 하나씩 빼냄
	//     - 출력값이 맨 앞에서 발생
	//       (1) 배열
	//         - 요소를 빼낼 때 마다 O(N) 성능이 필요
	//       (2) 리스트
	//         - O(1)
	//     → 배열보다 리스트가 큐에 더 효율적임

	queue := NewQueue()
	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
	// 1 -> 2 -> 3 -> 4 ->
}
