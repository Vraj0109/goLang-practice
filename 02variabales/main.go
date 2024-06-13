package main

import "fmt"

var token string = "lkjj23lk4j"
const PrivateToken string = "Private2lk43lk23" //variable name starting with capital latter is considered as public variable
func main(){
	var username string
	fmt.Println(username)
	fmt.Printf("variable of type:%T\n",username)

	age := 12 //this can only be used inside of the method
	fmt.Println(age)
	fmt.Printf("variable of type:%T\n",age)


	fmt.Println(token)
	fmt.Printf("variable of type:%T\n",token)


	fmt.Println(PrivateToken)
	fmt.Printf("variable of type:%T\n",PrivateToken)
}