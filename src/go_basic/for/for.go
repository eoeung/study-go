package main

import "fmt"

func main() {
	// go by example
	i := 1
	// 1) while문 처럼, 조건식만 작성
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// 2) 기본 for문
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 3) 무한 루프
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// 무한 루프처럼 보이지만, break로 1번만 돌고 탈출
	for {
		fmt.Println("loop")
		break
	}

	// 예제로 배우는 GO 프로그래밍
	// 4) for - range
	names := []string{"홍길동", "이순신", "강감찬"}
	for index, name := range names {
		fmt.Println(index, name)
	}

	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // 5-1) continue
		}
		a++

		if a > 10 {
			break // 5-2) break
		}
	}
	if a == 11 {
		goto END // 5-3) goto LABEL
	}
	fmt.Println(a) // 출력되지 않는다. → goto END를 통해 바로 'END:'로 이동

END:
	fmt.Println("END")

	t := 0
L1:
	for {
		if t == 0 {
			break L1 // 5-4) break LABEL
			// (1) for문을 탈출
			// (2) L1 레이블로 이동
			// (3) break가 있는 for문을 건너뛰고 다음 코드 실행
		}
	}

	fmt.Println("OK")
}
