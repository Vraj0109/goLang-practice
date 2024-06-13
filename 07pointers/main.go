package main

import "fmt"

func main() {
	fmt.Println("pointers sections")

	myNumber := 23
	fmt.Println("this is the number: ",myNumber)
	var ptr = &myNumber;
	fmt.Println("value of actual pointer is: ",ptr);
	fmt.Println("value of actual value stored in aboves address: ",*ptr);

	*ptr = *ptr + 2;
	fmt.Println("now stored value in the myNumber is: ",myNumber);
}