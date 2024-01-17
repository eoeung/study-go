package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// [에러 랩핑]
// - 에러를 감싸서 새로운 에러를 생성

// Ex) 파일에서 텍스트를 읽어서 특정 타입의 데이터로 변환하는 경우
//   1) 파일 읽기에서 발생하는 에러
//   2) 텍스트의 몇 번째 줄의 몇 번째 칸에서 에러가 발생했는지 알면 더 유용함
//   → 파일 읽기에서 발생한 에러를 감싸고, 그 바깥에 줄과 칸 정보를 넣는다.

// 문자열에서 두 단어를 읽어서 각 숫자로 변환한 뒤 곱한 결과를 반환
func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) // 스캐너 생성
	scanner.Split(bufio.ScanWords)                      // 한 단어씩 끊어읽기 (한 글자가 아니라 한 단어씩 끊어 읽는다.)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pos:%d err:%w", pos, err) // 에러 감싸기
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pos:%d err:%w", pos, err) // 에러 감싸기
	}
	return a * b, nil
}

// 다음 단어를 읽어서 숫자로 변환하여 반환
// return 변환된 숫자, 읽은 글자 수, 에러
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { // 단어 읽기
		return 0, 0, fmt.Errorf("failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word) // 문자열을 숫자로 반환
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert word to int, word:%s err222:%w", word, err) // 에러 감싸기
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println("#####", err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { // 감싸진 에러가 NumError인지 확인
			fmt.Println("NumberError :", numError)
		}
		// 1) errors.As() 혹은 errors.Is()로 감싸진 에러를 처리할 수 있음
		// 2) func errors.As(err error, target any) bool
		//   (1) err's tree에서 target과 일치하는 오류를 발견
		//   (2) 발견되면 target을 해당 오류 값으로 설정, true 반환
	}
}

func main() {
	readEq("123 3") // 369

	readEq("123 abc")
	/*
		##### failed to readNextInt(), pos:4 err:failed to convert word to int, word:abc err222:strconv.Atoi: parsing "abc": invalid syntax
		NumberError : strconv.Atoi: parsing "abc": invalid syntax

		// 1) 'failed to readNextInt(), pos:4 err:' 까지 MultipleFromString()에서 발생한 에러감싸기 구문
		// 2) err: 뒤에 있는 'failed to convert word to int, word:abc err222:strconv.Atoi: parsing "abc": invalid syntax'는 readNextInt()에서 발생한 에러감싸기 구문
	*/
}
