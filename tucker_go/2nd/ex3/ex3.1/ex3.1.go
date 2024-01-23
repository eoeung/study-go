package main

// 함수명 + 설명을 주석으로 작성하는 것을 코딩 규약으로 권장하고 있음

// CalculateWage는 일한 시간과 성과에 따라 보수를 결정하는 함수
// workTime은 주 52시간을 넘을 수 없음
// successRate는 백분율로 0~100사이의 값
// 반환값은 만 원 단위 실수로 표시됨. Ex) 2.5일 경우 2만 5천원
func CalculateWage(workTime int, successRate float64) float64 {
	return float64(0)
}

// Go는 정적 컴파일 언어
// - Go 내부 환경 변수인 'GOOS', 'GOARCH'값을 변경해서 원하는 컴파일 파일을 생성할 수 있음

// $ go tool dist list
// - 가능한 운영체제와 아키텍쳐 목록을 볼 수 있음
