package main

import "fmt"

func main() {
	// 예제로 배우는 GO 프로그래밍
	// [슬라이스]
	// - 기본적으로는 배열에 기초하여 만들어짐
	// - 배열의 단점을 보완
	//     1) 고정된 크기를 미리 지정하지 않을 수 있음
	//     2) 크기를 동적으로 변경할 수 있음
	//     3) 부분 배열 발췌 가능

	// 1) 슬라이스 생성
	// (1) 배열 생성하는 것과 비슷한데, 크기는 지정해주면 안됨
	var tt = []int{1, 2, 3}        // 슬라이스 변수 선언
	fmt.Printf("%T, %v\n", tt, tt) // []int, [1 2 3]
	// var a [3]int // 배열이라면 크기를 지정해줬어야 함
	tt[1] = 100
	fmt.Println(tt)

	var slice1 []int = []int{}
	var slice2 []int = make([]int, 0, 0)

	fmt.Println(slice1) // []
	fmt.Println(slice2) // []

	// (2) make()사용
	// 슬라이스의 length, capacity를 지정할 수 있음
	sss := make([]int, 5, 10) // 슬라이스 타입, len, cap
	// zzz := make([3]int) // 첫번째 인자는 t Type인데, 타입은 slice, map, channel 중 선택해야함
	fmt.Println(len(sss), cap(sss)) // 5 10

	var xxx []int
	if xxx == nil {
		fmt.Println("Nil Slice")
		fmt.Println(len(xxx), cap(xxx)) // 0 0
	}

	// 2) 부분 슬라이스
	// 슬라이스[처음 인덱스 : 마지막 인덱스]
	// 처음 인덱스 : inclusive, 마지막 인덱스 : exclusive
	var yyy = []int{1, 2, 3, 4, 5}
	fmt.Println(yyy) // [1 2 3 4 5]]
	yyy = yyy[2:]
	fmt.Println(yyy) // [3 4 5]

	// 3) 슬라이스 추가, 병합, 복사
	// (1) 슬라이스 추가
	//     - 배열은 고정된 크기 이상으로 데이터를 추가할 수 없음
	// append(슬라이스, 추가할 값)
	// → 추가할 값에는 슬라이스 객체도 가능 // 다만, sliceA...으로 ellipsis 사용해야함
	// 해당하는 슬라이스 객체의 모든 요소의 집합을 의미
	kk := []int{0, 1}
	kk = append(kk, 2)
	kk = append(kk, 3, 4, 5, 6)
	fmt.Println(kk)

	// (2) 슬라이스 병합
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}
	sliceD := append(sliceA, sliceB...)
	// sliceE := append(sliceA, sliceB..., sliceC...) // Error : expected ')', found sliceC
	fmt.Println(sliceD) // [1 2 3 4 5 6]

	// (3) 슬라이스 복사
	// copy(원래 슬라이스, 복사해서 붙여넣을 슬라이스)
	source := []int{1, 2, 3}
	target := make([]int, len(source), cap(source)*2)
	copy(source, target)
	fmt.Println(target)                   // [0 0 0]
	fmt.Println(len(target), cap(target)) // 3 6

	fmt.Println("###### go by example ######")
	// go by example
	s := make([]string, 3)
	s2 := make([]string, 10)
	fmt.Println("s", s)
	fmt.Println("s2", s2)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)      // [a b c]
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	// append는 새로운 값이 추가된 슬라이스를 반환
	s = append(s, "d")
	fmt.Println(s) // [a b c d]
	s = append(s, "e", "f")
	fmt.Println("append:", s) // [a b c d e f]

	c := make([]string, len(s))
	fmt.Println(c) // [     ]
	copy(c, s)
	fmt.Println("cpy:", c) // [a b c d e f]

	l := s[2:5]            // s[2], s[3], s[4]
	fmt.Println("sl1:", l) // [c d e]

	l = s[:5]
	fmt.Println("sl3:", l) // [a b c d e]

	l = s[2:]
	fmt.Println("sl4:", l) // [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // [g h i]

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD) // [[0] [1 2] [2 3 4]]
}
