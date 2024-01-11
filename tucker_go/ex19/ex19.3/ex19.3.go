package main

// [객체지향 프로그래밍]
// - 메서드라는 기능이 생기면서, 데이터와 기능을 묶을 수 있게 됨
// - 객체 간 관계가 매우 중요하다.

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

type Subject struct{}
type Professor struct{}
type Report struct{}

func (s *Student) EnrollClass(c *Subject) {
	// ...
}

func (s *Student) SendReport(p *Professor, r *Report) {
	// ...
}

// EnrollClass()와 SendReport()메서드의 리시버로 Student 구조체가 명시되어 있음
// → 두 메서드는 Student 구조체에 속함

// Student는 단순히 이름/나이 정보가 있는 데이터가 아니라,
// 과목을 등록하고, 리포트를 보낼 수 있는 기능이 추가된 객체가 됨
// → Subject, Professor, Report 모두와 관계를 맺는다!

// 함수 호출 순서보다는
// 객체를 만들고, 다른 객체와 상호관계를 맺는 것이 더 중요함
// → 객체 간 상호관계는 메서드로 표현됨!!
