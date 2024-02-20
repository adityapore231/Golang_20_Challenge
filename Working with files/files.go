package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("./test.txt")

	if err != nil {
		panic(err)
	}

	lenght, err := file.WriteString("I am doing practicle of golang")

	if err != nil {
		panic(err)
	}
	fmt.Println("lenght is ", lenght)
	defer file.Close()
}
