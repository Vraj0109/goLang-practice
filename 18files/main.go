package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files Module")

	textContent := "this is the sentence which will be written in the file"

	file, err := os.Create("./18files/newFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, textContent)

	if err != nil {
		panic(err)
	}
	fmt.Println("size is", length)

	readFile("./18files/newFile.txt")

	defer file.Close() //always use defer key if you want to keep writng after this line
}

func readFile(file string) {
	databyte, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}
	fmt.Println("the content inside the file in bytes is\n", databyte, "the content in string formate is\n", string(databyte))

}
