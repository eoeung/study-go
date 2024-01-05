package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	// [switch문에서 초기문 사용하기]

	// 1) 초기문 사용
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}

	// fmt.Println("age is", age) // Error : undefined: age → 지역변수 age는 switch문 안에서만 사용 가능

	// 2) 비교값을 true로 설정
	switch age := getMyAge(); true { // → switch age := getMyAge(); {...}로 사용 가능
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
}
