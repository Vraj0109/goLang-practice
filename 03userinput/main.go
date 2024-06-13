package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcomemsg := "hello from vraj"
	fmt.Println(welcomemsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza")

	// comma ok // err ok syntax

	input, _ := reader.ReadString('\n') // input, err for the error handling  _ is used when you don't care about it you can also use the _,err if you don't need the input for some reason(you sick idiot)
	fmt.Println("thanks for rating,",input)
}