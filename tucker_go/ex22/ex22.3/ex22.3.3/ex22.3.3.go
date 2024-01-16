package main

import "fmt"

func main() {
	// [맵에서 요소 삭제와 요소 확인]
	// 1) 요소 삭제
	//   - delete(map, key)
	//     - delete() 함수로 요소 삭제
	//     - 삭제 후, key값으로 요소값을 확인할 때 기본값으로 들어감
	//     - 그러면 값이 기본값(int일 경우라고 가정 → 0)일 때와 요소가 없을 때(기본값 반환) 모두 0을 반환한다.
	//       → 복수 반환으로 확인

	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	delete(m, 4) // 없는 요소 삭제 시도

	fmt.Println(m[3])
	fmt.Println(m[1])

	// 2) 요소 확인
	//   - 맵은 반환값을 하나 혹은 둘로 받을 수 있음
	//     (1) 반환값이 하나 : 요소값 반환
	//     (2) 반환값이 둘 : 요소값 + 요소 존재 여부 bool
	v, ok := m[1]
	fmt.Println(v, ok) // 0 true

	v, ok = m[3]
	fmt.Println(v, ok) // 0 false

	v, ok = m[4]
	fmt.Println(v, ok) // 0 false
}
