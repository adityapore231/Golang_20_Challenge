package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Sending get request to server")

	//sending get request to server
	response, err := http.Get("http://Localhost:8080")
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	bytecount, _ := responseString.Write(content)

	fmt.Println("Response is ", responseString.String())
	fmt.Println("Byte count is ", bytecount)

	// fmt.Println("Response is ", string(content))

}
