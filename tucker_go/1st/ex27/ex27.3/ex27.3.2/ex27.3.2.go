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

// Ex) 리스코프 치환 원칙 위배 예시

// 1) 함수를 생성시
// 2) 함수를 호출하는 호출자와 함수 구현 간의 계약 관계가 발생한다고 볼 수 있음

type Report interface {
	Report() string
}

// SendReport()의 의도를 추측
// - Report를 전송할 것으로 예상
//   → 호출자와 함수 간 계약이 성립
// 만약, 함수를 호출했는데 Report가 전송되지 않으면 예상치 못한 버그 발생
func SendReport(r Report) {
	// Report를 전송할 거라고 예상하는데
	if v, ok := r.(*MarketingReport); ok { // MarketingReport이 들어오면 패닉 발생
		fmt.Println(v, ok)
		panic("Can't send MarketingReport")
	}
}

type MarketingReport struct {
	name string
}

func (m *MarketingReport) Report() string {
	return fmt.Sprintf("%s 마케팅 보고서", m.name)
}

func main() {
	var report = &MarketingReport{name: "우리 회사"}
	SendReport(report)
}
