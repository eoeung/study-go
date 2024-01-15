package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3) // 함수 리터럴 3개를 가진 슬라이스
	fmt.Println("CaptureLoop")

	for i := 0; i < 3; i++ {
		fmt.Printf("## %p, %d\n", &i, i)
		/*
			## 0xc000096068, 0
			## 0xc000096068, 1
			## 0xc000096068, 2
		*/

		f[i] = func() {
			fmt.Printf("## f[i] :: %p :: %d\n", &i, i)
			// 1) 함수 리터럴 외부의 변수 i를 캡쳐할 때, 변수 i의 주소값을 포인터 형태로 함수 리터럴 내부 상태로 가져옴
			// 2) 나중에 캡쳐된 내부 상태를 사용할 때, 메모리 주소값을 통해 외부 변수 i에 접근
			/*
				## f[i] :: 0xc000096068 :: 3
				## f[i] :: 0xc000096068 :: 3
				## f[i] :: 0xc000096068 :: 3
				// → 위에 있는 for문의 지역변수 i를, 참조로 캡쳐한 것을 알 수 있음
			*/
			// fmt.Println(i) // 캡쳐된 i값 출력
			// 1) i 변수를 캡쳐할 때, 캡쳐하는 순간의 i값이 복사되는 것이 아니라 i 변수가 참조로 캡쳐됨
			// 2) for문이 진행될 때 마다 i값이 증가해서 3이 되고
		}
	}

	fmt.Println("------------")

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("CaptureLoop2")

	for i := 0; i < 3; i++ {
		v := i // v 변수에 i값 복사
		// 1) v 변수는 for문 내부에서 선언됐음
		// 2) 매 루프마다 새로운 v 변수가 생성됨
		fmt.Printf("## i :: %p, %d\n", &i, i)
		fmt.Printf("## v :: %p, %d\n\n", &v, v)
		/*
			## i :: 0xc0000960a8, 0
			## v :: 0xc0000960b0, 0

			## i :: 0xc0000960a8, 1
			## v :: 0xc0000960b8, 1

			## i :: 0xc0000960a8, 2
			## v :: 0xc0000960c0, 2
			// i는 동일하지만, v는 계속 다른 메모리 주소인 것을 알 수 있음
		*/

		f[i] = func() {
			fmt.Printf("## f[i] :: %p :: %d\n", &i, i)
			fmt.Printf("## f[v] :: %p :: %d\n\n", &v, v)
			// fmt.Println(v) // 캡쳐된 v값 출력

			/*
				## f[i] :: 0xc0000960a8 :: 3
				## f[v] :: 0xc0000960b0 :: 0

				## f[i] :: 0xc0000960a8 :: 3
				## f[v] :: 0xc0000960b8 :: 1

				## f[i] :: 0xc0000960a8 :: 3
				## f[v] :: 0xc0000960c0 :: 2
				// 1) i, v 모두 for문 지역 변수 i와 v를 가리키는 것을 알 수 있음
			*/
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	// [캡쳐]
	// - 함수 리터럴 외부 변수를 내부 상태로 가져오는 것
	// - 값 복사가 아닌 참조 형태로 가지고 온다.

	CaptureLoop()
	/*
		CaptureLoop
		## 0xc00000a0c8, 0
		## 0xc00000a0c8, 1
		## 0xc00000a0c8, 2
		------------
		## f[i] :: 0xc00000a0c8 :: 3
		## f[i] :: 0xc00000a0c8 :: 3
		## f[i] :: 0xc00000a0c8 :: 3
	*/

	fmt.Println()

	CaptureLoop2()
	/*
		CaptureLoop2
		0
		1
		2
	*/
}
