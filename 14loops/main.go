package main

import "fmt"

func main() {
	fmt.Println("Loops module");

	days := []string{"sun","mon","tue","wed","thu","fri","sat"};


	fmt.Println("with normal forloop syntax");
	for d:= 0;d < len(days);d++{
		fmt.Println("day : ",days[d]);
	}
	fmt.Println("syntax to use in array");
	for i := range days{
		fmt.Println("with:",days[i]);
	}

	for index, day := range days{
		fmt.Println("index " ,index," value: ",day);
	}

	fmt.Println("while loop");

	cnt := 1;

	for cnt < 10{
		fmt.Println("value is ",cnt);
		cnt++;
	}
}