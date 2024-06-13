package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type course struct { // it starts with small case so it's not exported
	Name      string   `json:"courseName"`
	Price     int      `json:"coursePrice"`
	Plateform string   `json:"plateform"`
	Password  string   `json:"-"`                    // this field will not be shown
	Tags      []string `json:"courseTags,omitempty"` // to handle null and don't write it as , omitempty
}
func main() {
	fmt.Println("Web reqests verbs module")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3001"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	println("Status code : ", res.StatusCode)
	println("Content Length : ", res.ContentLength)

	dataByte, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataByte)
	println("Content is: ", content)

	// other way of converting bytes to readable data in golang
	var responceString strings.Builder

	byteCount, _ := responceString.Write(dataByte)

	fmt.Println("byteCount is : ", byteCount)

	fmt.Println("Actual daya with strings.builder is : ", responceString.String())

	defer res.Body.Close()
}

func PerformPostJsonRequest() {

	fmt.Println("\n")
	const myurl = "http://localhost:3001/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"name":"vraj",
			"age":"10"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	dataByte, _ := io.ReadAll(res.Body)

	content := string(dataByte)

	fmt.Println("the content is : ", content)

	defer res.Body.Close()
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3001/postform"
	fmt.Println("\n")

	data := url.Values{}
	data.Add("firstName", "Vraj")
	data.Add("lastName", "Limbachiya")
	data.Add("Collage", "Indian Institute of information technology surat")
	data.Add("CGPA", "8")

	res, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	dataByte, _ := io.ReadAll(res.Body)

	content := string(dataByte)

	fmt.Println("Content from the formData is :", content)
}
