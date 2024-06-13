package main

import "fmt"

func main() {
	fmt.Println("Defer module")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	myDefer();

}

func myDefer(){
	for i := 1;i < 5;i++{
		defer fmt.Println(i)
	}
}