package main

import "fmt"

func main() {
	// [break, fallthrough]

	// 1) break
	//   - Go에서는 타언어와 다르게 break를 써도, 안써도 하나의 case를 실행하고 switch문 종료
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	case 3:
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}

	// a == 3

	// 2) fallthrough
	//   - 하나의 case문 실행한 다음, 다음 case문 까지 실행하고 싶은 경우

	b := 3

	switch b {
	case 1:
		fmt.Println("b == 1")
		break
	case 2:
		fmt.Println("b == 2")
	case 3:
		fmt.Println("b == 3")
		fallthrough
	case 4:
		fmt.Println("b == 4")
	case 5:
		fmt.Println("b == 5")
	default:
		fmt.Println("b > 5")
	}

	// b == 3
	// b == 4
}
