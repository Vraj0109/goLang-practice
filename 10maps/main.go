package main

import "fmt"
func main(){
	fmt.Println("The maps section");

	lang := make(map[string]string);
	lang["javascript"] = "gujarati";
	lang["python"] = "English";
	lang["c++"] = "chinese";
	lang["golang"] = "japanese";
	fmt.Println("The map of language is: ",lang );

	delete(lang,"python"); // delete using the key;
	fmt.Println("The map of language is: ",lang );

	//loop through the map

	for key,value := range lang{                // you can use the _,value or key,_ if you dont want to use the other other variable
		fmt.Println("key",key,":",value);
	}

}