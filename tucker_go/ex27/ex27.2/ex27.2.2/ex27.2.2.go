package main

// [객체 지향 설계 원칙 SOLID]
// 1) Single Responsibility (단일 책임 원칙)
// 2) Open-Closed (개방-폐쇄 원칙)
// 3) Liskov Substitution (리스코프 치환 원칙)
// 4) Interface Segregation (인터페이스 분리 원칙)
// 5) Dependency Inversion (의존 관계 역전 원칙)

// 2) Open-Closed
//   - 확장에는 열려있고, 변경에는 닫혀있다.
//   - 상호 결합도를 줄여서 새 기능을 추가할 때, 기존 구현을 변경하지 않아도 됨
//   → 프로그램에 기능을 추가할 때 기존 코드의 변경을 최소화해야 한다.

// Ex) 개방-폐쇄 원칙을 위반 → 올바른 예시

// 보고서
type Report struct {
	report string
}

// 전송 인터페이스 생성
type ReportSender interface {
	Send(r *Report)
}

// 이메일 전송 객체
type EmailSender struct {
}

// 이메일 전송
func (e *EmailSender) Send(r *Report) {
	// ...
}

// 팩스 전송 객체
type FaxSender struct {
}

// 팩스 전송
func (f *FaxSender) Send(r *Report) {
	// ...
}
