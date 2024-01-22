package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	// [if 초기문; 조건문]
	// - if문을 검사하기 전, 초기문 작성할 수 있음
	// - 초기문은 검사에 사요할 변수를 초기화할 때 사용
	// - 초기문에서 선언한 변수의 범위는, if문 안에서만 한정적으로 사용 가능

	if age, ok := getMyAge(); age < 20 && ok {
		fmt.Println("You are young", age)
	} else if age < 30 && ok {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	// fmt.Println("Your age is ", age) // Error : undefined: age
	// → if초기문에 선언된 변수는, if문이 종료되기 전까지 유지됨
}
