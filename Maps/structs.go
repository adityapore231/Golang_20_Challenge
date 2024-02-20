package main

import "fmt"

//structs in go

func main() {
	fmt.Println("Structs in Go")
	aditya := Employee{"aditya", "pore", 25}
	fmt.Println(aditya) //{aditya pore 25}
}

type Employee struct {
	firstName string
	lastName  string
	age       int
}
