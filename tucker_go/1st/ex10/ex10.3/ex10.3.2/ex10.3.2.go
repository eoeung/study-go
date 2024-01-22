package main

import "fmt"

func main() {
	// [switch문의 다양한 형태]

	// 2) 조건문 비교
	//   - if문 처럼 true도 비교할 수 있음

	temp := 18

	switch true { // → switch {...}로 줄여서 사용 가능
	case temp < 10, temp > 30: // → temp < 10 || temp > 30
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니, 가벼운 겉옷을 준비하세요.")
	// 위의 case문을 실행했기 때문에, 아래의 case문을 실행하지 않고 switch문은 종료됨
	case temp >= 15 && temp < 25:
		fmt.Println("야외 활동하기 좋은 날씨입니다.")
	default:
		fmt.Println("따뜻합니다.")
	}

	// 약간 추울 수 있으니, 가벼운 겉옷을 준비하세요.
	// → 2번째 case문을 만족하므로, 실행하고 switch문 종료
}
