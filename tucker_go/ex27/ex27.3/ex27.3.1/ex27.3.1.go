package main

import "fmt"

// [객체 지향 설계 원칙 SOLID]
// 1) Single Responsibility (단일 책임 원칙)
// 2) Open-Closed (개방-폐쇄 원칙)
// 3) Liskov Substitution (리스코프 치환 원칙)
// 4) Interface Segregation (인터페이스 분리 원칙)
// 5) Dependency Inversion (의존 관계 역전 원칙)

// 3) Liskov Substitution
//   (1) q(x)를 타입 T의 객체 x에 대해 증명할 수 있는 속성이라고 가정
//   (2) S가 T의 하위 타입
//   (3) q(y)는 타입 S의 객체 y에 대해 증명할 수 있어야 한다.

// ∴ 쉽게 풀어쓰기
// 1) 인터페이스 T(Something() 메서드)를 생성
// 2) S와 U 구조체가 Something() 메서드를 구현 → 인터페이스 T를 구현
// 3) q() 함수는 인터페이스 T를 인자로 받는다.
// 4) 객체 S의 인스턴스인 y와 객체 U의 인스턴스 u를 생성
// 5) y와 u는 인터페이스 T를 구현했으므로, q() 함수의 인자인 인터페이스 T로 전달 가능
// 6) q(y), q(u)는 문제없이 작동

type T interface { // SomeThing() 메서드를 포함한 인터페이스
	Something()
}

type S struct {
}

func (s *S) Something() { // T 인터페이스 구현

}

type U struct {
}

func (u *U) Something() { // T 인터페이스 구현

}

func q(t T) {
	fmt.Printf("### %T\n", t)
}

var y = &S{} // S 타입 y
var u = &U{} // U 타입 u

func main() {
	// 둘 다 잘 작동해야함
	q(y)
	q(u)
}
