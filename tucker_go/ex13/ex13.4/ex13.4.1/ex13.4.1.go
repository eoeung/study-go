package main

import "fmt"

type Student struct {
	Age   int
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이:%d 번호:%d 점수:%.2f\n", s.Age, s.No, s.Score)
}

func main() {
	// [구조체 크기]
	// - 구조체 필드를 모두 담을 수 있도록, 필드를 연속되게 담을 수 있는 메모리 공간 할당

	// 1) 구조체 값 복사
	var student = Student{15, 23, 88.2}
	student2 := student // student 구조체의 모든 필드가 student2로 복사됨

	PrintStudent(student2) // 함수 호출 시에도 구조체가 복사됨
}
