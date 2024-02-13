package main

import "fmt"

//maps in go

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)
	languages["en"] = "English"
	languages["fr"] = "French"
	languages["es"] = "Spanish"
	languages["de"] = "German"

	fmt.Println(languages) //map[de:German en:English es:Spanish fr:French]

	delete(languages, "fr")
	fmt.Println(languages) //map[de:German en:English es:Spanish]

	for key, value := range languages {
		fmt.Println(key, value)
	}

	//list of format specifiers available in go are:
	// %v - value in default format
	// %+v - value in default format with field names
	// %#v - value in go syntax
	// %T - type of the value
	// %t - true or false
	// %d - base 10
	// %b - base 2
	// %c - character
	// %x - base 16
	// %f - floating point
	// %e - scientific notation
	// %s - string
	// %q - quoted string
	// %p - pointer
	// %x - base 16
	// %X - base 16 with upper case letters
	// %U - unicode
	// %q - double quoted string

}
