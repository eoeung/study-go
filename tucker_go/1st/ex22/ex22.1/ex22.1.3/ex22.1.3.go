package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{v: list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	// [스택 구현하기]
	// - FILO

	// - 스택의 특징
	//   1) 가장 최근에 넣은 것부터 역순으로 나옴 → 마지막에 넣은 것이 제일 먼저 나옴
	//   2) 요소는 맨 뒤로 추가
	//   3) 요소를 뺄 때도 맨 뒤에서 빼냄

	// - 함수 호출에 스택을 사용
	//   - a() → b() → c() 호출하면, c() 종료되면 b() → a()로 돌아가야 하기 때문
	//   - 호출 순서와 역순으로 진행되어야 함

	stack := NewStack()
	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}
	// 4 -> 3 -> 2 -> 1 ->
}
