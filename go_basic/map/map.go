package main

import "fmt"

func main() {
	// 예제로 배우는 GO 프로그래밍
	// [Map]
	// map[Key type]value type
	// key : 정수 타입 / value : 문자열

	// 1. Map 선언
	// 1) var 사용
	var idMap map[int]string

	// 2) make() 사용
	idMap = make(map[int]string)
	// make 순서
	// (1) 해시테이블 자료구조를 메모리에 생성
	// (2) 이 메모리를 가리키는 map value를 리턴
	//       - runtime.hmap 구조체를 가리키는 포인터
	// (3) idMap 변수 = 이 해시테이블을 가리키는 map을 가리킴

	if len(idMap) == 0 {
		fmt.Printf("%T, %v\n", idMap, idMap) // map[int]string, map[]
	}

	// 3) Literal 사용해서 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}
	fmt.Printf("%T, %v\n", tickers, tickers) // map[string]string, map[FB:FaceBook GOOG:Google Inc MSFT:Microsoft]

	// 2. Map 사용
	// Error : assignment to nil map
	// var m map[int]string
	// m[901] = "Apple"
	// m[134] = "Grape"
	// m[777] = "Tomato"

	m := make(map[int]string)
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	fmt.Println(str)

	noData := m[999]                       // 값이 없으면 nil or zero or "" 리턴
	fmt.Printf("%T, %v\n", noData, noData) // string, ""

	// 3. Map 키 체크
	val, exists := tickers["MSFT"]
	fmt.Println(val, exists) // Microsoft true
	if !exists {
		fmt.Println("No MSFT ticker")
	}

	// 4. Map 열거
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range로 모든 맵 요소 출력
	// Map은 unordered이므로 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}

	// go by example
	ttt := make(map[string]int)
	ttt["k1"] = 7
	ttt["k2"] = 13
	fmt.Println(ttt)      // map[k1:7 k2:13]
	fmt.Println(len(ttt)) // 2 // len(Map)은 Map에 있는 key,value 쌍의 개수 반환

	v1 := ttt["k1"]
	fmt.Println(v1) // 7

	// delete(Map, key)
	// Map의 key,value 삭제
	delete(ttt, "k2")
	fmt.Println(ttt) // map[k1:7]

	// Map에서 값을 꺼내오면, (val, key 존재 여부) 반환
	value, exist := ttt["k1"]
	fmt.Println(value, exist) // 7 true
	value, exist = ttt["kkkkkk"]
	fmt.Println(value, exist) // 0 false

	// 한 줄로 Map 선언
	n := map[string]int{"foo": 1, "bar": 2} // map[bar:2 foo:1]
	fmt.Println(n)

	var sss map[int]int
	var zzz map[int]int = map[int]int{1: 1}
	fmt.Println(sss, zzz) // map[] map[1:1]

	// Map의 Key값만 순회
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// 1) key-value 모두 출력
	for k, v := range kvs {
		fmt.Println(k, v)
	}
	/*
		a apple
		b banana
	*/

	// 2) key만 출력
	for k := range kvs {
		fmt.Println(k)
	}
	/*
		a
		b
	*/
}
