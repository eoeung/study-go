package main

func main() {
	// [리터럴]
	// - 리터럴 : 고정된 값 (값 자체로 쓰인 문구)
	// - Go에서는 상수를 리터럴 취급 → 컴파일시 리터럴로 변환됨
	//   → 상수의 메모리 주소에 접근할 수 없는 이유 : 리터럴로 전환되어 실행파일에 값으로 쓰이기 때문

	var str string = "Hello World"
	var i int = 0
	i = 30

	// 상수 표현식에서는 컴파일 타임에서 실제 리터럴로 변환함
	// 아래 같은 코드가 있을 때, 최종적으로는 18 Line 코드로 컴파일됨
	const PI = 3.14
	var a int = PI * 100

	// 14 ~ 15 Line 코드 컴파일시
	var a int = 314
}
