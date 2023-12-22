package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	dbDialect := os.Getenv("DB_DIALECT")
	dbInfo := os.Getenv("DB_INFO")
	updateVideoNm := os.Getenv("UPDATE_VIDEONM")

	db, err := sql.Open(dbDialect, dbInfo)
	checkError(err)
	defer db.Close()

	// UPDATE문 실행
	// 1) Exec() 사용
	//   - DML같이 리턴값이 없는 경우, Exec() 사용
	result, err := db.Exec(updateVideoNm, "videonm")
	checkError(err)

	// RowsAffected() : 갱신된 레코드 수 반환
	// LastInsertId() : 새로 추가된 id 반환
	cnt, _ := result.RowsAffected()
	fmt.Printf("%d rows updated\n", cnt)

	// 2) Prepared Statement
	stmt, err := db.Prepare(updateVideoNm)
	checkError(err)
	_, err = stmt.Exec("videonn")
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
