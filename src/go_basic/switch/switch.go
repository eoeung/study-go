package main

import (
	"fmt"
	"time"
)

// 2) 함수의 인자로 작성
func grade(score int) {
	switch {
	case score >= 90:
		fmt.Println("A입니다.")
	default:
		fmt.Println("F입니다.")
	}
}

// fallthrough
func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
	case 4:
		fmt.Println("4 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}

func main() {
	i := 2
	fmt.Println(i)

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other number")
	}

	// case 문에 복잡한 표현식 적용
	switch i {
	case i % 2:
		fmt.Println("홀수")
	default:
		fmt.Println("짝수")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch에 바로 중괄호 사용 가능
	// switch {...} 라고 작성하면, switch expression을 true라고 생각하고 case 문으로 이동
	// 1) 미리 선언한 변수 값으로 작성
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	// 2) 함수의 인자로 작성
	grade(99)
	grade(50)

	// 변수 타입 검사
	whatAmI := func(i any) {
		// whatAmI := func(i interface{}) {
		s, ok := i.(string)
		fmt.Println(i)
		fmt.Printf("s : %v, ok : %v, \n", s, ok) // s는 string에 한해서만 일치하는 값이 나온다.

		switch i.(type) { // type assertion
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("???")
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(1.5)

	// Go에서는 break가 있든 없든 관계없이 항상 break를 진행함
	// → Go 컴파일러가 자동으로 각 case문 블럭 마지막에 break를 추가
	// fallthrough : 타 언어처럼 break 사용
	check(2)
}
