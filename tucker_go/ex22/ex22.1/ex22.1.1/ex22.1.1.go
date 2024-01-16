package main

import (
	"container/list"
	"fmt"
)

type Element struct { // 리스트의 각 요소 데이터를 저장하는 구조체
	Value interface{} // 데이터를 저장하는 필드
	Next  *Element    // 다음 요소의 주소를 저장하는 필드
	Prev  *Element    // 이전 요소의 주소를 저장하는 필드
}

func main() {
	// [리스트]
	// - 불연속 메모리에 데이터를 저장
	// - 각 데이터를 담고 있는 요소를, 포인터로 연결한 자료구조

	v := list.New()       // 새로운 리스트 생성
	e4 := v.PushBack(4)   // 리스트 뒤에 요소 추가
	e1 := v.PushFront(1)  // 리스트 앞에 요소 추가
	v.InsertBefore(3, e4) // e4 요소 앞에 요소 삽입
	v.InsertAfter(2, e1)  // e1 요소 뒤에 요소 삽입

	for e := v.Front(); e != nil; e = e.Next() { // 각 요소 순회
		fmt.Print(e.Value, " ") // 1 2 3 4
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() { // 각 요소 역순 순회
		fmt.Print(e.Value, " ") // 4 3 2 1
	}
}
