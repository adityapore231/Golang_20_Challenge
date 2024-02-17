package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("hello world")
	fmt.Println("This is a defer tutorial.")

	defer fmt.Println("hello world 2") //printing in reverse order
	defer fmt.Println("hello world 3")

	print_defer()

}

//when use defer keyword, the function call is deferred to the end of the function.

func print_defer() {
	for i := 0; i < 2; i++ {
		defer fmt.Println(i)
	}
}
