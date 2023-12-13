package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	// path := `D:\download\go.png` // `` 사용하면 escape 하지 않음
	path := "D:/download/go.png"
	file, _ := os.Open(path) // 경로에 있는 파일을 읽음
	defer file.Close()       // defer : 함수가 종료되기 직전에 실행

	os.RemoveAll("./uploads")

	// Web으로 파일 전송할 때는 MIME 포맷을 사용해야함
	buf := &bytes.Buffer{}             // NewWriter()에 넣을 버퍼 생성 → 데이터가 버퍼에 저장됨
	writer := multipart.NewWriter(buf) // MIME 전송하기 위해 multipart.NewWriter()사용
	// 폼 파일 생성
	// → 데이터 넣어줘야함
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file) // io.Copy(destination, original)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)          // 저장된 데이터가 있는 버퍼를 서버에 전송
	req.Header.Set("Content-type", writer.FormDataContentType()) // 폼 데이터 타입임을 명시해줌

	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	// 실제로 같은 파일인지 확인
	uploadFilePath := "./uploads/" + filepath.Base(path)
	fileInfo, err := os.Stat(uploadFilePath)
	assert.NoError(err)
	fmt.Println(fileInfo)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	fmt.Println(uploadFilePath, uploadFile)
	fmt.Println(path, originFile)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := make([]byte, 102400)
	originData := make([]byte, 102400)
	uploadLen, _ := uploadFile.Read(uploadData)
	originLen, _ := originFile.Read(originData)

	fmt.Println(uploadLen, originLen)
	assert.Equal(uploadData, originData)
}
