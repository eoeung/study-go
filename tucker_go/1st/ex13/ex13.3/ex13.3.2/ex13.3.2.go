package main

import "fmt"

// [구조체를 포함하는 구조체]
//   1) 내장 타입처럼 포함하는 방식
//   2) 포함된 필드 방식

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User     // 2) 포함된 필드 방식 → 필드명 생략
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만 원\n",
		vip.Name, // vip.UserInfo.Name처럼 작성하지 않고, .하나로 접근할 수 있음
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price, // 닫는 소괄호 ')'가 마지막 인수와 다른 줄에 있으면 ',' 작성해야 함
	)
}
