package main

import "fmt"

func main() {

	//var ptr *int
	//fmt.Println(ptr)
	myval := 23
	var ptr = &myval
	fmt.Println(ptr)  //0xc00000a0a8 address
	fmt.Println(*ptr) //23 value
	*ptr = *ptr * 2
	fmt.Println(myval) //46
}
