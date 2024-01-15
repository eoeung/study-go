package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) { // Writer 함수 타입, 변수 호출
	writer("Hello World")
}

func main() {
	// [파일 핸들을 내부 상태로 사용할 때]

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer f.Close()

	writeHello(func(msg string) { // 파일에 msg를 쓰는 함수 리터럴 생성 → writeHello() 함수의 인수로 사용
		fmt.Fprintln(f, msg)
		// 1) 함수 리터럴 외부 변수인 f 사용
		// 2) Hello World 문자열을 인수로 호출
	})
	// → 의존성 주입
	//   - 외부에서 로직을 주입하는 것
	//   - writeHello()는 인수로 온 writer를 호출했을 때, 어떤 로직이 수행되는지 알 수 없음
	//     Ex) 파일에 씌여지거나, 네트워크에 전송되거나, 프린터로 찍히거나 어떤 것이 수행되는지 알 수가 없음
}
