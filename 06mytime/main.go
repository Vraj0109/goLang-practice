package main

import (
	"fmt"
	"time"
)

func main(){

	section := "mytime"

	fmt.Println(section);

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}