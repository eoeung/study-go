// 핸들러를 만들어주는 패키지
package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// UpdateUser struct
type UpdateUser struct {
	ID               int       `json:"id"`
	UpdatedFirstName bool      `json:"updated_first_name"`
	FirstName        string    `json:"first_name"`
	UpdatedLastName  bool      `json:"updated_last_name"`
	LastName         string    `json:"last_name"`
	UpdatedEmail     bool      `json:"updated_email"`
	Email            string    `json:"email"`
	CreatedAt        time.Time `json:"created_at"`
}

var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	// 유저가 없는 경우
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}

	users := []*User{}
	for _, u := range userMap {
		users = append(users, u)
	}
	data, _ := json.Marshal(users)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

//	func getUserInfo89Handler(w http.ResponseWriter, r *http.Request) {
//		fmt.Println(r.URL.Path) // /users/89
//		fmt.Fprint(w, "User Id:89")
//	}
func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.Path)
	// vars := mux.Vars(r) // map[id:555555]
	// fmt.Println(vars)
	// fmt.Fprint(w, "User Id:", vars["id"])

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func deleteUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	_, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", id)
		return
	}

	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ID:", id)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	updateUser := new(UpdateUser) // 업데이트 원하는 유저 정보
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user, ok := userMap[updateUser.ID]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", updateUser.ID)
		return
	}
	fmt.Println(userMap[updateUser.ID])

	// userMap[updateUser.ID] = updateUser
	// ※ 주의점. 요청을 보낼 때, 모든 key값이 아닌 수정하고 싶은 key값만 보내는 경우,
	//    key값 없이 온 정보 값은 해당 key값의 default값이 들어옴
	// → 근데, 그걸 모두 덮어쓰기 해버림
	if updateUser.UpdatedFirstName {
		user.FirstName = updateUser.FirstName
	}
	if updateUser.UpdatedLastName {
		user.LastName = updateUser.LastName
	}
	if updateUser.UpdatedEmail {
		user.Email = updateUser.Email
	}
	// userMap[updateUser.ID] = user // 할 필요가 없음
	// → user, ok := userMap[updateUser.ID] 코드에서 이미 userMap은 포인터를 반환
	// => 위의 if문에서 user가 이미 해당 객체를 가리키므로

	fmt.Println(userMap[updateUser.ID])
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Println(string(data))
	fmt.Fprint(w, string(data))
}

// myapp 핸들러 생성
func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	lastID = 0
	mux := mux.NewRouter() // gorilla mux 사용

	// mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")
	// mux.HandleFunc("/users/89", getUserInfo89Handler) // 하드코딩
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserInfoHandler).Methods("DELETE")
	return mux
}
