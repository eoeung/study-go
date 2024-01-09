package main

import "fmt"

func main() {
	// [문자열]
	// - Go에서는 UTF-8 문자코드 사용
	// - 큰 따옴표(") / 백쿼트(`)

	// 2) 여러 줄 문자열 처리 가능 여부
	poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를,\n잎새에 이는 바람에도\n나는 괴로워했다.\n"

	poet2 := `죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를,
잎새에 이는 바람에도
나는 괴로워했다.`

	fmt.Println(poet1)
	fmt.Println(poet2)
}
