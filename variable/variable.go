package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// global variabel

const auth_token = "aajshdkjdhsiodhaso6876"

func main() {
	var username string = "aditya"
	fmt.Println(username)
	fmt.Printf("type of varible is %T \n", username)

	var onoff bool = true
	fmt.Println(onoff)
	fmt.Printf("type of varible is %T \n", onoff)

	var smallval uint8 = 245
	fmt.Println(smallval)
	fmt.Printf("type of varible is %T \n", smallval)

	var dummy uint8
	fmt.Println(dummy)
	fmt.Printf("type of varible is %T \n", dummy)

	var floatval float32 = 245.4654576876
	fmt.Println(floatval)
	fmt.Printf("type of varible is %T \n", floatval)

	//245.46545

	var floatval64 float64 = 245.4654576876
	fmt.Println(floatval64)
	fmt.Printf("type of varible is %T \n", floatval64)

	// we can declare varible without mentiosning it's type
	var str = "mynameisaditya"
	fmt.Println(str)

	// withouth using var keyword := is called as walrus operator
	a := 898
	fmt.Println(a)

	// use global variable
	fmt.Println(auth_token)

	// taking input from user
	welcome_msg := "welcome"
	println(welcome_msg)

	fmt.Println("enter your age")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	fmt.Println("Your age is", input)

	//handling Input and conversion
	age, _ := strconv.ParseUint(strings.TrimSpace(input), 10, 8)
	if err != nil {
		fmt.Println("Error parsing:", err)
	} else {
		fmt.Println("Parsed value:", age+2) // Output: 123
	}
}
