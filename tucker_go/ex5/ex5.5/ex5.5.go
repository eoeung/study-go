package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Scan()함수는 호출되면 새로운 입력을 받는 것이 아니라, 기존에 남아있는 표준 입력 스트림에서 값을 가져온다.
	// → 여러 번 Scan()함수 호출할 때 표준 입력 스트림에 있던 값으로 진행하게 된다.
	// => 기존에 남아있는 표준 입력 스트림을 비워주면 이런 에러가 발생하지 않음

	stdin := bufio.NewReader(os.Stdin) // 표준 입력을 읽는 객체

	var a int
	var b int

	// 첫 번째 입력받기
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)

		str, _ := stdin.ReadString('\n') // 표준 입력 스트림 지우기 → \n이 나올 때 까지 읽으면 표준 입력 스트림이 비워짐
		fmt.Println(str)
	} else {
		fmt.Println(n, a, b)
	}

	// 두 번째 입력받기
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}

	/*
		Hello 4
		0 expected integer
		ello 4 // 기존에 남아있는 표준 입력 스트림에 있던 값들

		3 4
		2 3 4
	*/
}
