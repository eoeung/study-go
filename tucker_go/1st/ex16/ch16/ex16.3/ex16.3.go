package main

import (
	"ch16/ex16.3/exinit"
	"fmt"
)

// [패키지 초기화]
// 1) 패키지를 import하면, 컴파일러는 패키지 내 전역변수 초기화함
// 2) init()함수가 있으면 호출해서 패키지를 초기화
//   - init()은 입력 매개변수, 반환값이 없어야함
//   - 어떤 패키지의 초기화 함수인 init()만 실행하고 싶은 경우, _을 이용해 import
// → 전역 변수가 모두 초기화 되면, init()함수 호출

func main() {
	fmt.Println("main function")
	exinit.PrintD()

	/*
		f() d: 4
		f() d: 5
		init function 6
		main function
		d: 6
	*/
}
