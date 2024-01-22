package main

import "fmt"

func main() {
	score := 85
	count := 15

	switch {
	case count < 10:
		fmt.Println("평가 수가 모자랍니다.")
	case count < 20 && score > 80:
		fmt.Println("긍정적인 평가")
	case count < 20:
		fmt.Println("판단할 단계가 아닙니다.")
	case score > 90:
		fmt.Println("좋은 평가입니다.")
	case score > 80:
		fmt.Println("살 만한 물건입니다.")
	case score > 70:
		fmt.Println("신중히 생각해보세요.")
	default:
		fmt.Println("좋은 물건이 아닙니다.")
	}

	// 긍정적인 평가
}
