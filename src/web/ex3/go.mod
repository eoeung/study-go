module goproject/ex3

go 1.21.5

replace goproject/ex3/myapp => ./myapp // goproject/ex3/myapp : main.go에서 사용(참조)할 패키지명 / ./myapp : 실제 로컬에 있는 app.go 경로
