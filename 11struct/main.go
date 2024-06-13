package main

import "fmt"

func main() {
	fmt.Println("Struct Module")

	vraj := User{"vraj", "vrajlimbachiya99@gmail.com", true, 21}

	fmt.Println("The current user is: ", vraj)

	fmt.Printf("Vraj's details are : %+v\n", vraj) //batter way to print struct with formating
}

// user uppercase

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
