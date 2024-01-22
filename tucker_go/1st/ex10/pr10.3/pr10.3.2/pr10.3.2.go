package main

import "fmt"

// 책에 있는 알고리즘
//   - const 열거값을 사용해서 작성

type Direction int

const (
	None  Direction = iota // 0
	North                  // 1
	East                   // 2
	South                  // 3
	West                   // 4
)

func DirectionToString(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func GetDirection(angle float64) Direction {
	switch {
	case angle >= 315 || (angle >= 0 && angle < 45):
		return North
	case angle >= 45 && angle < 135:
		return East
	case angle >= 135 && angle < 225:
		return South
	case angle >= 225 && angle < 315:
		return West
	default:
		return None
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

	fmt.Println(DirectionToString(GetDirection(38.3)))
	fmt.Println(DirectionToString(GetDirection(235.8)))
	fmt.Println(DirectionToString(GetDirection(94.2)))
	fmt.Println(DirectionToString(GetDirection(-30)))
	fmt.Println(DirectionToString(GetDirection(135)))
}
