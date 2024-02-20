package main

import "fmt"

func main() {

	fmt.Println("This is a struct tutorial.")

	type user struct {
		name     string
		email    string
		age      int
		employed bool
		family   []string
	}

	user1 := user{"aditya", "adityapore001@gmail.com", 21, true, []string{"aditya", "rohini", "sandeep"}}

	fmt.Println(user1.family[2])
	fmt.Printf("Name is %v and email is %v", user1.name, user1.email)
}
