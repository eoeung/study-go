package main

import "fmt"

// recover() 함수 선언 → panic()과 동일하다.
// func recover() interface{}

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		if r := recover(); r != nil { // recover()는 발생한 panic()객체를 반환 / 없으면 nil
			fmt.Println("panic 복구 -", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func main() {
	// [패닉 전파 그리고 복수]
	// - 프로그램이 종료되는 대신, 에러 메시지를 표시하고 복구를 시도하는 것이 더 좋을 수도 있음
	// - main() 함수에서도 복구가 안되면 프로그램 강제 종료
	// - 패닉은 복구된 시점부터 프로그램이 계속됨

	// - recover() 함수가 호출되는 시점에 패닉이 전파 중이면, panic객체 반환 / 아니면 nil 반환
	f()
	fmt.Println("프로그램이 계속 실행됨")

	// recover()는 조심해서 사용해야 함
	//   - 복구가 되더라도 프로그램이 불안정할 수 있음
	//     Ex) 파일에 데이터를 쓰고있다가 패닉이 발생한 경우
	//       → 복구하더라도 데이터는 일부만 써진 상태임
	//       => 복구하지 않는 등 데이터를 비정상적으로 남아있지 않게 하는 등 작업이 필요
}
