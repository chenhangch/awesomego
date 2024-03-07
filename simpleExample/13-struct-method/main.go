package main

import "fmt"

type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}
func main() {
	a := user{
		name:     "chen",
		password: "100",
	}
	var d user
	d.name = "chen"
	d.password = "100"

	d.checkPassword("100")
	d.resetPassword("2048")

	fmt.Println(a, d)
}


