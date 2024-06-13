package main

import "fmt"

func main() {
	fmt.Println("Functions module")
	result := add(10, 5)
	fmt.Println("result is ", result)

	proResult, name := proAdder("vraj", 10, 23, 34, 132, 232)
	fmt.Println("pro result is ", proResult, " name is ", name)
	hello()
}

func proAdder(s string, values ...int) (int, string) {
	result := 0
	for _, value := range values {
		result += value
	}
	return result, s
}
func add(a int, b int) int {
	return a + b
}
func hello() {
	fmt.Println("Hello wolrd!")
}
