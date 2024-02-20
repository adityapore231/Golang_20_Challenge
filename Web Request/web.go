package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/ashishps1/awesome-system-design-resources?tab=readme-ov-file"

func main() {
	fmt.Println("Handling web requests...")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	//fmt.Println("Response is ", response)
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("Data is ", string(data))
	defer response.Body.Close()

}
