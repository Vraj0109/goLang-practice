package main

import "fmt"

func main() {
	fmt.Println("Methods Module")

	vraj := User{"vraj", 21, "vrajlimbachiy99@gmail.com", true}

	fmt.Println("user is ", vraj)
	fmt.Printf("user is %+v\n", vraj) // batter way of printing the struct

	userName := vraj.GetName();

	fmt.Println("user's name is ", userName)

}

type User struct {
	Name   string
	Age    int16
	Email  string
	Status bool
}

func (u User) GetName() (User) {
	u.Name = "someone";
	return u;
}
