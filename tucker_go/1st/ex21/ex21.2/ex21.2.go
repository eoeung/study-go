package main

import (
	"fmt"
	"os"
)

func main() {
	// [defer 지연 실행]
	// - OS 내부 자원 사용하는 파일, 소켓 핸들 같이 함수가 종료되기 직전에 실행해야하는 코드
	//   1) 파일을 생성하거나 읽을 때, OS에 파일 핸들을 요청
	//   2) OS는 파일 핸들을 만들어서 프로그램에 알려줌
	//   3) OS 내부 자원이기 때문에 쓰고 나서 반드시 OS에 되돌려줘야함
	//     → 프로그램에서 OS 내부 자원을 되돌려주지 못하면, 내부 자원이 고갈되어 파일을 못 만들거나 통신이 불가능할 수도 있음
	//     => 프로그램이 종료되면 자동으로 모든 자원을 반환하지만, 실행 중에는 프로그램에서 직접 반환해줘야 함

	// defer
	// - 함수가 종료되기 직전에 실행되도록 지연됨

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	// 지연 수행될 코드
	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World") // 파일에 텍스트 작성

	/*
		파일에 Hello World를 씁니다.
		파일을 닫았습니다.
		반드시 호출됩니다.
	*/
	// 호출 순서가 'defer의 역순'으로 호출됨
}
