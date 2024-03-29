package main

import "fmt"

const M = 10 // 나머지 연산의 분모

func hash(d int) int {
	return d % M // 나머지 연산
}

func main() {
	// [해시 함수]
	// 1) 같은 입력값 → 같은 결과값
	// 2) 다른 입력값 → 되도록 다른 결과값
	// 3) 입력값의 범위는 무한대지만, 결과값은 특정 범위 내

	// Ex-1) 삼각 함수
	//   1) sin(90) = 1
	//   2) sin(90), sin(450), sin(-90), sin(-450) = 1 이지만, 일반적으로 다른 결과값이 나옴
	//   3) 입력값은 무한대지만, 결과값은 -1 ~ 1 사이로 나옴

	// → 삼각 함수는 계산이 복잡하고 결과값이 실수로 나오기 때문에 해시 함수로 사용 X
	// => 일반적으로 나머지 연산을 사용

	// Ex-2) 나머지 연산
	//   → f(x) = Mod(x, 10)
	//     1) f(5), f(15), f(25) = 5
	//     2) 입력값이 다르면 왠만하면 다른값을 가짐, 10마다 같은 결과값을 가짐
	//     3) 입력값은 무한대이지만, 결과값은 0 ~ 9 사이

	// 1) 나머지 연산은 계산이 간단해 매우 빠름
	// 2) 결과값의 범위와 간격을 조절하기 쉬움

	/////////////////////////////////////////////////////
	// [해시로 맵 만들기]
	// 1) 해시 함수는 결과값이 항상 일정한 범위(개수)를 가짐
	// 2) 같은 입력 → 같은 결과 보장
	// 3) 일정 범위에서 반복
	// → 범위와 같은 요소 개수를 갖는 배열이 적합

	m := [M]int{} // 값을 저장할 배열 생성 → 해시 함수 결과값은 (0 ~ M-1) => M개

	m[hash(23)] = 10 // m[idx] = value → 배열이니깐
	m[hash(33)] = 10 // 해시 충돌 → 해시 함수가 다른 입력값임에도 같은 결과나 나옴
	// → hash(23)과 hash(33) 모두 인덱스 위치가 3이다.
	m[hash(259)] = 50

	fmt.Printf("%d = %d\n", 23, m[hash(23)])
	fmt.Printf("%d = %d\n", 259, m[hash(259)])

	// [해시 충돌을 방지하려면]
	// - 인덱스 위치마다 값이 아닌 리스트를 저장

	// ∴ Go에서 map은 해시 함수를 통해 key를 해시값으로 변환한다.
	//     1) 각 key의 해시값은 고유하다.
	//     2) 이 해시값으로 데이터를 저장하고 검색한다.
}
