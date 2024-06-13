package main

import "fmt"

func main(){
	fmt.Println("Array module");

	var arr [4]string;

	arr[0] = "my";
	arr[1] = "name";
	arr[3] = "vraj";
	arr[2] = "";

	fmt.Println("arr is: ",arr);
	fmt.Println("size of the array is",len(arr))

}