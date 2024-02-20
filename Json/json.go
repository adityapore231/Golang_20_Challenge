package main

import (
	"encoding/json"
	"fmt"
)

type laptop struct {
	Name      string `json:"Model Name"` // We can create alies for this data . Very Imp : Not to put space in between
	Ram       int    `json:"Ram in GB"`
	Rom       int
	Processor string `json:"-"` //hide content
	Price     float32
	Tags      []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("This is about json handling")
	DecodeJson()

}

func ConvertoJson() []byte {
	Dell := laptop{"Dell", 3, 32, "Intel", 32000, []string{"office user", "General"}}

	//finaljson, _ := json.Marshal(Dell)
	finaljson, _ := json.MarshalIndent(Dell, "", "\t")

	//fmt.Printf("%s\n", finaljson)
	return finaljson
}

func DecodeJson() {
	json_data := ConvertoJson()
	var data laptop
	checkValid := json.Valid(json_data)
	if checkValid {
		json.Unmarshal(json_data, &data)
		fmt.Printf("%#v\n", data)
	} else {
		fmt.Println("Json is not valid")
	}

	var laptopdata map[string]interface{} //we can store it as a map
	json.Unmarshal(json_data, &laptopdata)
	fmt.Printf("%#v\n %T\n ", laptopdata, laptopdata)

	//fmt.Println(string(json))
}
