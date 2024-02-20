package main

import "fmt"

func main() {
	fmt.Println(add(4, 6))

	fmt.Println(adder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func add(a int, b int) int {
	return a + b
}

func adder(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
