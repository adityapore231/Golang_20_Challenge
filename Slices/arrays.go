package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to arrys")

	var arr [4]string
	arr[0] = "banana"
	arr[1] = "apple"
	arr[3] = "pear"
	fmt.Println("fruit list is ", arr)

	var vegetables = [3]string{"potato", "tomato", "cucumber"}
	for i := 0; i < len(vegetables); i++ {
		fmt.Println(vegetables[i])
	}
	fmt.Print(len(vegetables)) // output: 3
}
