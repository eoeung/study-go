package main

func main() {
	// [중첩 for문과 break, 레이블]
	//   - break는 속해있는 for문만 빠져나온다.
	//   - 모든 for문을 나가고 싶을 때, 사용하는 방법

	// 2) 레이블 사용
	var a, b int = 1, 1

OuterFor: // 레이블 정의 → '레이블명:'으로 정의
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor // 레이블에 가장 먼저 포함된 for문까지 종료
			}
		}
	}
}
