package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices Module");

	var arr = []string{"apple","watermelone"}
	fmt.Printf("type of arr is %T\n",arr);
	fmt.Println("Element of the array are:",arr);

	arr = append(arr, "mango","banana");
	fmt.Println("updated Element of the array are:",arr);

	fmt.Println("Slice the array from index 1 : ",arr[1:3]); // [startIndex,endIndex+1];

	highScores := make([]int, 4);

	highScores[0] = 232;
	highScores[1] = 332;
	highScores[2] = 821;
	highScores[3] = 10;

	// highScores[4] = 203;  //you cant use this since the initial allocatoin of the memory is of 4 elements

	highScores = append(highScores, 23,123,53); //but you can do this cause append will re-alocate the memory to the variable;

	sort.Ints(highScores )// to sort the slices
	

	fmt.Println("value of highscore:" ,highScores);

	//delete the element of certain index;

	var lang = []string{"c++","python","javascript","ruby","golang","java"};

	fmt.Println("initial array: ",lang)

	var index int = 2;
	lang = append(lang[:index], lang[index+1:]... ) //learn why to use ... spred oprator here
	fmt.Println("the 2nd index is deleted : ",lang)
}