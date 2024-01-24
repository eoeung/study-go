package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// [표준 입력 스트림]
	// 1) 표준 입력 장치로 값을 입력
	// 2) 표준 입력 스트림이라는 메모리 공간에 임시로 저장됨

	// 3-1) Scan()에서 에러가 발생하지 않으면
	//   → 표준 입력 스트림 값을 모두 읽기 때문에, 표준 입력 스트림은 비어있는 상태임

	// 3-2) Scan()에서 에러가 발생하면
	//   → 표준 입력 스트림은 에러가 발생한 상황 그대로 남아있게 되는 상태임 => 값이 남아있음
	//   → 추후 다시 Scan()을 시도하면, 사용자가 입력한 값이 아니라 남아있는 값 부터 읽게 됨

	// Ex) Scanln()이 실패했을 때, 표준 입력 스트림을 비워주는 예제
	stdin := bufio.NewReader(os.Stdin) // 표준 입력을 읽는 객체

	var a, b int
	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n')
		// '\n'은 rune 타입(int32)지만, ReadString(delim byte)은 byte 타입인자를 가진다.
		// 에러가 발생하지 않는 이유는 '\n'은 ascii code 10번(LF)이기 때문에 byte 타입으로 나타낼 수 있어서 가능하다.
		// → byte는 uint8 타입 (0 ~ 255) => '\n' = 10
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

	fmt.Printf("%T :: %v\n", '\n', '\n')
	var newline interface{} = '\n'
	if _, ok := newline.(byte); ok {
		if _, ok := newline.(rune); ok {
			fmt.Println("byte & rune")
		} else {
			fmt.Println("byte")
		}
	} else {
		fmt.Println("no byte")
	}
}
