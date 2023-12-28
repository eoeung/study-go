package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	// input에서 보낸 파일을 읽어오는 부분
	uploadFile, header, err := r.FormFile("upload_file") // key = input의 id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	fmt.Println(uploadFile)
	// fmt.Println(header)

	// 저장할 파일을 만드는 부분
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777) // 폴더가 없으면 새로 만들어줌 // 0777은 권한
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err2 := os.Create(filepath) // 새로 만들었지만 비어있는 file을 반환
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	defer file.Close() // err2 확인 전에 이 코드가 삽입되면, err2를 확인하라는 경고 등장

	// 받은 파일을 아까 만들었지만 비어있는 file에 복사
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	// 파일에 접근할 수 있는 권한을 public경로에 있는 파일에 부여?
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/uploads", uploadsHandler)

	http.ListenAndServe(":3000", nil)
}
