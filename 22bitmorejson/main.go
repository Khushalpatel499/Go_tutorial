package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Learn myjson in go")
	//EncodedJson()
	DecoderJson()
}

//encoded: covert data into valid json

func EncodedJson() {
	lcoCourse := []course{
		{"ReactJs Bootcamp", 299, "learnCoderonlin.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "learnCoderonlin.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "learnCoderonlin.in", "hit123", nil},
	}
	// package this data as json data

	//finalJson, err := json.Marshal(lcoCourse)

	// finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")

	finalJson, err := json.MarshalIndent(lcoCourse, "lco", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecoderJson() {

	jsonDataFromWeb := []byte(`
	        {
		             "coursename": "ReactJs Bootcamp",
		            "Price": 299,
		            "tags": ["web-dev","js"]
		    }
	`)
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json is not valid")
	}
	//in some cases you want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("The key is :%v and values is :%v and type is:%T\n", k, v, v)
	}
}
