package model

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3" // 명시적으로 사용하지는 않지만, database/sql이 아닌 이 패키지를 사용함을 의미
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos(sessionId string) []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT id, name, completed, createdAt FROM todos WHERE sessionId = ?", sessionId) // 세션에 해당하는 값만 가지고 와야함
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() { // rows.Next() : true되어 무한 루프 → 더 이상 조회된 값이 없으면 false
		todo := Todo{}
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}

	return todos
}

func (s *sqliteHandler) AddTodo(sessionId string, name string) *Todo {
	stmt, err := s.db.Prepare("INSERT INTO todos (sessionId, name, completed, createdAt) VALUES (?, ?, ?, datetime('now'))")
	if err != nil {
		panic(err)
	}

	rst, err := stmt.Exec(sessionId, name, false) // 위에서 작성한 ?에 들어갈 값을 적어주면 됨
	if err != nil {
		panic(err)
	}
	id, _ := rst.LastInsertId()
	todo := Todo{}
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()
	return &todo
}

func (s *sqliteHandler) RemoveTodo(id int) bool {
	stmt, err := s.db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		panic(err)
	}

	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}

	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	stmt, err := s.db.Prepare("UPDATE todos SET completed = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}

	rst, err := stmt.Exec(complete, id)
	if err != nil {
		panic(err)
	}

	cnt, _ := rst.RowsAffected() // 영향 받은 row 수 반환
	return cnt > 0
}

// 추후 database를 종료하기 위해 만든 메서드
func (s *sqliteHandler) Close() {
	s.db.Close()
}

func newSqliteHandler(filepath string) DBHandler {
	// DB 생성/오픈
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	// todos 테이블 생성
	// 세션별로 모두 같은 데이터를 다루는 것이 문제 → 세션별 데이터 저장 (sessionId 컬럼 추가)
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id		  INTEGER PRIMARY KEY AUTOINCREMENT,
			sessionId STRING,
			name      TEXT,
			completed BOOLEAN,
			createdAt DATETIME
		);
		CREATE INDEX IF NOT EXISTS sessionIdIndexOnTodos ON todos(
			sessionId ASC
		);`)
	// → WHERE절에 있는 sessionId컬럼에 대한 인덱스를 생성

	// 쿼리 실행
	statement.Exec()

	// DB를 계속 사용해야하기 때문에, 값으로 저장해줌
	return &sqliteHandler{db: database}
}
