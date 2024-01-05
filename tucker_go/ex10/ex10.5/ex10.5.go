package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor(color ColorType) ColorType {
	return color
}

func main() {
	// [const 열거값 + switch]

	// Ex) 색을 나타내는 열거값을 문자열로 바꾸는 함수
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor(Red)))
}
