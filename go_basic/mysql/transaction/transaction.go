package main

import (
	"database/sql"
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

	db, err := sql.Open(dbDialect, dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 트랜잭션 시작
	// 트랜잭션 : 복수 개의 SQL문 실행
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback() // 중간에 에러 발생하면, Rollback

	// UPDATE 문 실행
	tx.Exec("") // Exec()는 반환할 값이 없는 쿼리문 작성 할 때, 사용

	// 트랜잭션 커밋
	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}
