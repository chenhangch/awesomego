package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{
		name:     "chen",
		password: "100",
	}
	var d user
	d.name = "chen"
	d.password = "100"

	fmt.Println(a, d)
	fmt.Println(checkPassword(a, "haha"))
	fmt.Println(checkPassword2(&a, "100"))
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}