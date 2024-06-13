package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://dummy.restapiexample.com/api/v1/employees"

func main() {
	fmt.Println("Webrequests Module")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("responce type is%T\n", res)

	defer res.Body.Close()

	databytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println("data is ", content)
}
