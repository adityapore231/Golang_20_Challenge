package main

import "fmt"

func main() {
	var veglist = []string{"Tomato", "Onion", "Cucumber"}
	fmt.Println(veglist) //[Tomato Onion Cucumber]

	veglist = append(veglist, "bringle", "potato")
	fmt.Println(veglist) //[Tomato Onion Cucumber bringle potato]

	veglist = append(veglist[1:3])
	fmt.Println(veglist) //[Onion Cucumber]

	age := make([]int, 4)
	age[0] = 23
	age[1] = 25
	age[2] = 26
	//age[3] = 27

	fmt.Print(age)                                       // [22 26 23 25]
	fmt.Println("\nLength of age array is : ", len(age)) // Length of age array is :  4

	// Accessing out of range index will cause a runtime error
}
