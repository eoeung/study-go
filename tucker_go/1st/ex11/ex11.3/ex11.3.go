package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// [continue와 break]
	// 1) continue
	//   - 후처리로 건너뜀

	// 2) break
	//   - for문을 종료

	// Ex) 짝수를 입력하면 종료되는 프로그램
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요.")

		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요.")

			// 키보드 버퍼 지워줌
			stdin.ReadString('\n') // 인자로 받은 값이 처음 발생할 때까지 읽음
			continue
		}

		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 {
			break
		}
	}

	fmt.Println("for문이 종료됐습니다.")
}
