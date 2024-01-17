package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) // 파일 열기
	if err != nil {
		return "", err
	}
	defer file.Close() // 함수 종료 직전, 파일 닫기

	rd := bufio.NewReader(file)    // 파일 내용 읽기
	line, _ := rd.ReadString('\n') // error는 delim 문자로 안 끝나는 경우에 반환됨 → 무시
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) // 파일 생성
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "data.txt"

func main() {
	// [에러 핸들링]
	// - 적절한 방식으로 에러를 처리하는 것이 매우 중요

	// Ex) 파일 읽기에 실패하면 임시 파일 생성하고, 다시 한 번 파일 읽기 시도
	line, err := ReadFile(filename) // 파일 읽기 시도
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") // 파일 생성
		if err != nil {                                // 에러 처리
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}

		line, err = ReadFile(filename) // 다시 읽기 시도
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}

	fmt.Println("파일내용 :", line)
}
