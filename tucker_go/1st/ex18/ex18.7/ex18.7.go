package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// []Student의 별칭 타입 Students
type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age } // Less니까 i가 작은지 비교
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	// [슬라이스 정렬]
	// - sort 패키지로 슬라이스 정렬

	// 1) int 슬라이스 정렬
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)

	// 2) float64 슬라이스 정렬
	sss := []float64{2.2, 3.4, 0.2, -9.2, 1.1, -30.2}
	sort.Float64s(sss)
	fmt.Println(sss)

	// 3) 구조체 슬라이스 정렬
	// - Sort()함수를 이용하기 위해서는, Len(), Less(), Swap() 메서드 필요
	// → 이 메서드만 구현해서 커스터마이징
	sd := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42},
		{"켄", 38}, {"송하나", 18},
	}

	sort.Sort(Students(sd)) // sort.Interface를 사용할 수 있음
	// 1) Students(sd)는 []Student타입인 sd를 정렬 인터페이스를 포함한 타입인 Students 타입으로 변환
	// 2) []Student 타입은 Len(), Less(), Swap() 메서드가 없기 때문에 sort.Sort()함수의 인자로 사용될 수 없음
	// 3) []Student의 별칭 타입인 Students를 만들어서 정렬 인터페이스를 포함하도록 설정

	fmt.Println(sd) // [{송하나 18} {화랑 31} {켄 38} {류 42} {백두산 52}]

	/*
		1) Sort() 함수는 Interface를 인자로 받는다.
			func Sort(data Interface) {...}
		2) Interface라는 interface는 Len(), Less(), Swap()을 구현해주면 됨
			type Interface interface {
				Len() int
				Less(i, j int) bool
				Swap(i, j int)
			}
	*/
}
