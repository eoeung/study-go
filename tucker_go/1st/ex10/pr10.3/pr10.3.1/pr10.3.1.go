package main

import "fmt"

// 나의 알고리즘
// - 원시적으로 값을 비교하는 방법

type Direction string

func GetDirection(angle float64) Direction {
	switch {
	case (angle >= 0 && angle < 45) || angle >= 315:
		return "North"
	case angle >= 45 && angle < 135:
		return "East"
	case angle > 135 && angle < 225:
		return "South"
	case angle >= 225 && angle < 315:
		return "West"
	default:
		return "None"
	}
}

func main() {
	// [조건]
	// 함수명은 GetDirection
	// 함수 매개변수는 angle float64를 받음
	// 함수 결과는 Direction 타입 반환

	// [세부적인 조건]
	// 1) angle이 315 이상이면, North 반환
	// 2) angle이 0 이상 45 미만이면 North 반환
	// 3) angle이 45 이상 135 미만이면 East 반환
	// 4) angle이 135 이상 225 미만이면 South 반환
	// 5) angle이 225 이상 315 미만이면 West 반환
	// 6) 모든 조건을 만족하지 않으면 None 반환
	angle := 1111
	fmt.Println(GetDirection(float64(angle)))
}
