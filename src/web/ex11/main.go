package main

import (
	"html/template" // escape 처리 // Ex) eoeung113%40gmail.com
	"os"
	// "text/template" // escape 처리 안함 // Ex) eoeung113@gmail.com
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "moon", Email: "eoeung113@gmail.com", Age: 23}
	user2 := User{Name: "aaa", Email: "aaa@gmail.com", Age: 40}
	users := []User{user, user2}

	// tmpl, err := template.New("Tmpl1").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge: {{.Age}}")
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}

	// Parse() 사용시
	// tmpl.Execute(os.Stdout, user)
	// tmpl.Execute(os.Stdout, user2)

	// ParseFiles() 사용시
	// tmp1l.tmpl
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user)
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user2)

	// // tmpl2.tmpl
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user)
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)

	// users
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
