<div align="center">

## 💬 About This Repository

<p>
<h3>기초 Go 스터디</h3>
</p>

<br>

## 🔨 Stacks

<div>
    <!-- Language -->
    <p><strong>Language</strong></p>
    <div>
        <img src="https://img.shields.io/badge/Go-007d9c?style=for-the-badge&logo=Go&logoColor=ffffff"> 
    </div>
</div>

<br>
<br>

## 🔍 <font color="green">__How to use__</font>
<div align="left">

1. Go 설치
    ```bash
    # go 버전 확인
    $ go version

    # go 환경변수 확인
    $ go env
    ```
    - 기본적으로 아래 경로에 설치됨 (Windows)
        > C:\Program Files\Go
    - GOPATH와 GOROOT가 있는데, 자세히 살펴봐야함
2. Git 설치

<br>

### | Hello World 출력하는 간단한 예제
1. workspace 폴더 생성
2. hello.go 파일 생성
3. 아래 코드를 작성
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello World")
    }
    ```
    > ※ VSCode로 진행할 때는, Go Extension을 설치해야함 <br> 설치 에러는 무시해도 됨
4. Terminal에서 아래와 같이 명령어 작성
    ```bash
    # 1) hello.go를 실행 → Hello World 출력
    $ go run hello.go

    # 2) hello.exe 생성 (windows)
    $ go build
    # 이 명령어를 실행하면 아래와 같은 에러 발생
    # go: go.mod file not found in current directory or any parent directory; see 'go help modules'

    # 3) go mod init 프로젝트명
    $ go mod init study-go/hello

    # 4) 생성된 hello.exe 파일 실행
    $ hello.exe
    ```

</div>