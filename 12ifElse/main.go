package main

import "fmt"

func main() {
	fmt.Println("ifElse Module")

	cnt := 2

	var ans string

	if cnt < 10 {
		ans = "cnt is less then 10"
	}
	fmt.Println(ans)

	// special syntax
	if a := 3; a < 10 {
		fmt.Println("the number is less then 10")
	} else {
		fmt.Println("the number is greater then or equal to 10")
	}
}
