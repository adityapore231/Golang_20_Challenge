package main

func main() {

	user1 := user{"aditya", "aditya@gmail", 21, true, true}
	user2 := user{"aditya", "aditya@gmail", 21, false, true}
	println(user1.isEmployed())
	println(user2.isEmployed())

}

type user struct {
	name     string
	email    string
	age      int
	employed bool
	status   bool
}

func (u user) isEmployed() bool {
	return u.employed
}
