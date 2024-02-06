package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	year, _, _ := presentTime.Date()
	fmt.Println(year)
	fmt.Println(presentTime.Year())
	fmt.Println(presentTime.Clock())

}
