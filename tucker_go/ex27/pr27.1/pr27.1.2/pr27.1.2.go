package main

import "fmt"

// 의존 관계 역전 원칙 적용해보기
// solution

type Attacker interface {
	Name() string
}

type Defender interface {
	DealDamage(attacker Attacker, damage int)
}

type Player struct {
	name string
}

type Monster struct {
	hp int
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Attack(defender Defender) {
	defender.DealDamage(p, 100)
}

func (m *Monster) DealDamage(attacker Attacker, damage int) {
	m.hp -= damage
	if m.hp < 0 {
		fmt.Println(attacker.Name(), "가 나를 죽였다.")
	}
}
