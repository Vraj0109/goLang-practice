package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case module")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("The dice number is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("the dice number is 1 so you can move 1 spot")
	case 2:
		fmt.Println("the dice number is 2 so you can move 2 spot")
	case 3:
		fmt.Println("the dice number is 3 so you can move 3 spot")
	case 4:
		fmt.Println("the dice number is 4 so you can move 4 spot")
	case 5:
		fmt.Println("the dice number is 5 so you can move 5 spot")
	default:
		fmt.Println("The dice value is 6 so you can open or move 6 spots and you'll get 1 extra turn")
	}
}
