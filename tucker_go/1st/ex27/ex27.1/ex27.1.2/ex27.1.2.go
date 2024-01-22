package main

import (
	"fmt"
	"time"
)

// [객체 지향 설계 원칙 SOLID]
// 1) Single Responsibility (단일 책임 원칙)
// 2) Open-Closed (개방-폐쇄 원칙)
// 3) Liskov Substitution (리스코프 치환 원칙)
// 4) Interface Segregation (인터페이스 분리 원칙)
// 5) Dependency Inversion (의존 관계 역전 원칙)

// 1) Single Responsibility
//   Ex) 단일 책임 원칙에 위배 예시를 올바르게

// Report() 메서드를 포함한 Report 인터페이스
type Report interface {
	Report() string
}

// 회계 보고서
type FinanceReport struct {
	report string
}

// Report 인터페이스 구현
func (f *FinanceReport) Report() string {
	return f.report
}

// 마케팅 보고서
type MarketingReport struct {
	name string
	date time.Time
}

func (m *MarketingReport) Report() string {
	return fmt.Sprintf("%s 마케팅 보고서 작성 시간 : %v\n", m.name, m.date)
}

// 보고서 전송 담당
type ReportSender struct {
}

// Report 인터페이스 객체를 인수로 받아서 회계, 마케팅 등 모두 사용 가능하게 설정
func (r *ReportSender) SendReport(report Report) { // 함수 인자를 인터페이스로 전달하기 때문에, 포인터로 전달해야 함
	// 1) 인터페이스는 메서드의 집합
	// 2) 인터페이스에는 메서드를 호출할 수 있는 값을 가지고 있어야 함
	// 3) 인터페이스를 값으로 호출하면 원본 객체에 영향을 미치지 않음
	receiveReport = report.Report()
}

var receiveReport string

func main() {
	// 만약 실제 사용한다면?
	// 회계 보고서 생성
	f := FinanceReport{report: "회계 보고서 123"}

	// 회계 보고서를 전송
	rs := ReportSender{}
	rs.SendReport(&f)
	fmt.Println(receiveReport)

	////////////////////////
	// 마케팅 보고서 생성
	m := MarketingReport{name: "우리 회사", date: time.Now()}

	// 마케팅 보고서 전송
	rs.SendReport(&m)
	fmt.Println(receiveReport)
}
