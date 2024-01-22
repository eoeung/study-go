package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	// [패닉]
	// - 프로그램을 정상 진행시키기 어려운 상황을 만났을 때, 프로그램 흐름을 중지
	// - 문제 발생 시점에 panic()을 사용해 프로그램을 바로 종료
	// - 콜 스택 표시
	//   - 콜 스택 : 함수 호출 순서
	//   -
	//   - 버그 수정에 유용함

	// Ex) 나눗셈 함수에서 제수가 0일 때 panic() 호출

	divide(9, 3)
	divide(9, 0)
	/*
		9 / 3 = 3
		panic: b는 0일 수 없습니다.

		goroutine 1 [running]:
		main.divide(0x9?, 0x3?)
				D:/workspace/src/study-go/tucker_go/ex23/ex23.4/ex23.4.go:7 +0xee
		main.main()
				D:/workspace/src/study-go/tucker_go/ex23/ex23.4/ex23.4.go:24 +0x29
		exit status 2
	*/
	// call stack : panic이 발생한 마지막 함수 위치부터 역순으로 호출 순서를 표시
}
