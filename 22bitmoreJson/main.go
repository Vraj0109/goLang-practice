package main

import (
	"encoding/json"
	"fmt"
)

type course struct { // it starts with small case so it's not exported
	Name      string   `json:"courseName"`
	Price     int      `json:"coursePrice"`
	Plateform string   `json:"plateform"`
	Password  string   `json:"-"`                    // this field will not be shown
	Tags      []string `json:"courseTags,omitempty"` // to handle null and don't write it as , omitempty
}

func main() {

	fmt.Println("BitMoreJson Module")

	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"reactJs", 299, "lernReact.in", "abcd1234", []string{"webdev", "js"}},
		{"golang", 299, "lernGolang.in", "password", []string{"webdev", "go"}},
		{"Nodejs", 299, "lernNode.in", "lksjdlfs", nil},
	}

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println("the json data is:", string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"courseName": "reactJs",
			"coursePrice": 299,
			"plateform": "lernReact.in",
			"courseTags": ["webdev","js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON DATA WAS NOT IN VALID FORMATE")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", k, v,v)
	}

}
