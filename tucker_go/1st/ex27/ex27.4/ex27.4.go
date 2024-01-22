package main

import "time"

// [객체 지향 설계 원칙 SOLID]
// 1) Single Responsibility (단일 책임 원칙)
// 2) Open-Closed (개방-폐쇄 원칙)
// 3) Liskov Substitution (리스코프 치환 원칙)
// 4) Interface Segregation (인터페이스 분리 원칙)
// 5) Dependency Inversion (의존 관계 역전 원칙)

// 4) Interface Segregation (인터페이스 분리 원칙)
//   - 클라이언트는 자신이 사용하지 않는 메서드에 의존하지 않아야 한다.

// Ex) 인터페이스 분리 원칙 위반 예시

type Report interface {
	Report() string
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}

// SendReport()는 r이 가진 메서드 중에 Report() 메서드만 사용
// → r 인터페이스 이용자에게 불필요한 메서드를 인터페이스가 포함하고 있음
// Pages(), Author(), WrittenDate() 3가지

// 위의 예시를 올바르게 수정하면

type Report interface {
	Report() string
}

type WrittenInfo interface {
	Pages() int
	Author() string
	WrittenInfo() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}
