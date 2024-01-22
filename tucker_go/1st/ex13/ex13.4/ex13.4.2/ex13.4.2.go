package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32
	Score float64
}

func main() {
	// [구조체 크기]
	// 2) 필드 배치 순서에 따른, 구조체 크기 변환

	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user)) // 16 → 12바이트가 아닌 16바이트가 출력됨 → '메모리 정렬' 때문
	// unsafe.Sizeof(): 해당 변수의 메모리 공간 크기를 반환

	// [메모리 정렬]
	// - 레지스터 4바이트 : 32비트 컴퓨터
	// - 레지스터 8바이트 : 64비트 컴퓨터
	// → 한 번 연산에 레지스터 크기만큼 연산이 가능
	// => 데이터가 레지스터 크기와 똑같은 크기로 정렬되어 있으면, 효율적으로 데이터를 읽어올 수 있음
}
