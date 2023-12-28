package main

import (
	"database/sql" // 사실상 개발자는 database/sql을 통해 모든 SQL 프로세싱을 진행
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // 개발자가 직접 드라이버 패키지를 사용하지 않게 진행 → 드라이버 패키지는 database/sql 패키지가 내부적으로 사용하게 됨
	"github.com/joho/godotenv"
)

type Menu struct {
	menuId int
	menuNm string
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	dbDialect := os.Getenv("DB_DIALECT")
	dbInfo := os.Getenv("DB_INFO")
	selectMenu := os.Getenv("SELECT_MENU")

	// 실제 DB Connection이 일어나지 않음 → 정보만 들고 있음
	// 실제로 DB연결이 필요한 시점에 Connection이 이루어짐
	db, err := sql.Open(dbDialect, dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // 모든 쿼리가 완료될 때 까지 대기한다.

	// [SELECT]
	// Single Row
	// QueryRow() 사용
	so := selectOne(db)
	fmt.Println("#### so ::: ", so)

	// Mutirows
	// db.Query() 사용
	sm := selectMany(db, selectMenu)

	for idx, menu := range sm {
		fmt.Printf("index :::, %d \t ##### menu ::: %v\n", idx, menu)
	}

}

// Single Row
func selectOne(db *sql.DB) int {
	var i int
	err := db.QueryRow("SELECT 1234 AS num").Scan(&i)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

// Multirows
func selectMany(db *sql.DB, query string) []*Menu {
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	menus := []*Menu{}
	for rows.Next() {
		menu := Menu{}
		// rows.Scan() : 현재 행의 열을, 대상이 가리키는 값으로 복사
		rows.Scan(&menu.menuId, &menu.menuNm)
		menus = append(menus, &menu)
	}

	return menus
}
