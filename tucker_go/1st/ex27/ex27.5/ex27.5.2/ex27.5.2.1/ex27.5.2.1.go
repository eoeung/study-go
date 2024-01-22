package main

// [객체 지향 설계 원칙 SOLID]
// 1) Single Responsibility (단일 책임 원칙)
// 2) Open-Closed (개방-폐쇄 원칙)
// 3) Liskov Substitution (리스코프 치환 원칙)
// 4) Interface Segregation (인터페이스 분리 원칙)
// 5) Dependency Inversion (의존 관계 역전 원칙)

// 5) Dependency Inversion (의존 관계 역전 원칙)
//   - 상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전(역전)시킴으로써
//     상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다.

//   원칙 2) 추상 모듈은 구체화된 모듈에 의존해서는 안된다.
//           구체화된 모듈은 추상 모듈에 의존해야한다.

// 원칙 2) 예제
// - 메일이 오면 알람이 울린다.

// 1) 의존 관계 원칙에 위배되는 예제
type Mail struct {
	alarm Alarm
}

type Alarm struct {
	Alarm()
}

func (m *Mail) OnRecv() { // OnRecv() 메서드는 메일 수신 시 호출됨
	m.alarm.Alarm() // 알람을 울림
}
// → 메일이라는 구체화된 모듈이 알람이라는 구체화된 모듈에 의존하고 있음