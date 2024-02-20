package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println("This is URL parsing tutorial")

	URL := "https://www.turing.com/interview-questions/kubernetes?tab=readme-ov-file#kubernetes-interview-questions"

	parsed_url, err := url.Parse(URL)

	if err != nil {
		panic(err)
	}

	fmt.Println("parserd_url is ", parsed_url)

	fmt.Println("Scheme is ", parsed_url.Scheme)
	fmt.Println("Host is ", parsed_url.Host)
	fmt.Println("Path is ", parsed_url.Path)
	fmt.Println("RawQuery is ", parsed_url.RawQuery)
	fmt.Println("Fragment is ", parsed_url.Fragment)
	fmt.Println("Port is ", parsed_url.Port())

	params := parsed_url.Query() // map[tab:[readme-ov-file]]
	fmt.Println(params["tab"])   // [readme-ov-file]

	//constructing url

	val := &url.URL{ //always pass url as reference
		Scheme:  "https",
		Host:    "www.turing.com",
		Path:    "/interview-questions/kubernetes",
		RawPath: "tab=readme-ov-file",
	}
	fmt.Printf("Type of constructed url is %T\n", val)
	fmt.Println("Constructed URL is ", val.String())
}

// what is fragment in URL?
// Fragment is the part of the URL that comes after the # symbol. It is used to point to a specific section of the page.
