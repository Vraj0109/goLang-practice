package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("welcome to te conversion module")
	fmt.Println("please rate out pizza")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n');
	fmt.Println("thanks for rating",input)

	newRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	

	if err != nil {

		fmt.Println("eror msg:",err)
	} else {
		fmt.Println("new Rating is:",newRating + 1)
	}
}