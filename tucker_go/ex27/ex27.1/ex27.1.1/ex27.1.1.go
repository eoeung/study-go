package main

// [객체 지향 설계 원칙 SOLID]
// 1) Single Responsibility (단일 책임 원칙)
// 2) Open-Closed (개방-폐쇄 원칙)
// 3) Liskov Substitution (리스코프 치환 원칙)
// 4) Interface Segregation (인터페이스 분리 원칙)
// 5) Dependency Inversion (의존 관계 역전 원칙)

// 1) Single Responsibility
//   - 모든 객체는 하나의 책임만 져야 한다.
//   Ex) 단일 책임 원칙에 위배 예시

// 회계 보고서
type FinanceReport struct {
	report string
}

// 보고서 전송
func (r *FinanceReport) SendReport(email string) {
	// ...
}

// → 회계 보고서라는 책임이 있는데, 전송이라는 책임도 있기 때문에 단일 책임 원칙에 위배

// 만약 마케팅 보고서라는 객체가 있을 경우
// 마케팅 보고서
type MarketingReport struct {
	report string
}

// 마케팅 보고서 전송
func (r *MarketingReport) SendReport(email string) {
	// ...
}

// → 보고서 전송 함수인 SendReport()를 사용할 수 없다.
// => *FinanceReport 타입의 메서드이기 때문

// 그렇다면, 단일 책임 원칙에 입각한 설계는?
// 1) FinanceReport → Report 인터페이스를 구현
// 2) ReportSender → Report 인터페이스를 이용하는 관계를 형성
