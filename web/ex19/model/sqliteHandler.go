package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // 명시적으로 사용하지는 않지만, database/sql이 아닌 이 패키지를 사용함을 의미
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) getTodos() []*Todo {
	return nil
}

func (s *sqliteHandler) addTodo(name string) *Todo {
	return nil
}

func (s *sqliteHandler) removeTodo(id int) bool {
	return false
}

func (s *sqliteHandler) completeTodo(id int, complete bool) bool {
	return false
}

// 추후 database를 종료하기 위해 만든 메서드
func (s *sqliteHandler) close() {
	s.db.Close()
}

func newSqliteHandler() dbHandler {
	// DB 생성/오픈
	database, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// todos 테이블 생성
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id		  INTEGER PRIMARY KEY AUTOINCREMENT,
			name      TEXT,
			completed BOOLEAN,
			createdAt DATETIME
		)`)
	// 쿼리 실행
	statement.Exec()

	// DB를 계속 사용해야하기 때문에, 값으로 저장해줌
	return &sqliteHandler{db: database}
}
