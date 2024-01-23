package main

import "fmt"

func main() {
	// [변수 선언]
	// - 컴퓨터에게 값을 저장할 공간을 마련하라고 명령을 내리는 것 → '메모리 할당'이라고 한다.

	var minimumWage int = 10
	var workingHour int = 20

	// 1) 정수 타입 데이터를 저장할 공간을 생성한다.
	// 2) 생성한 공간에 각각 minimumWage, workingHour라고 지칭한다.
	// 3) 각각 10과 20이라는 값을 복사한다.
	// → minimumWage와 workingHour라는 변수명을 통해서 해당 공간에 접근할 수 있음

	var income int = minimumWage * workingHour

	fmt.Println(minimumWage, workingHour, income)
}
