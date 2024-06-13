package main

import (
	"fmt"
	"net/url"
)


const myUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
const myUrl2 = "https://dummy.restapiexample.com/api/v1/employees"
func main() {
	fmt.Println("Ulr Module");

	fmt.Println(myUrl);
	fmt.Println(myUrl2);

	//parsing the url

	result,err := url.Parse(myUrl);

	if err != nil {
		panic(err)
	}
	fmt.Println("scheme is: ",result.Scheme)
	fmt.Println("host is: ",result.Host)
	fmt.Println("path is: ",result.Path)
	fmt.Println("port is: ",result.Port())
	fmt.Println("rawquery is: ",result.RawQuery)

	qparams := result.Query();

	fmt.Printf("the type of qparams is %T\n ",qparams)

	fmt.Println("the params is: ",qparams)

	for _,value := range qparams{
		fmt.Println("params is -> ",value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:"Vraj.com",
		Path: "/erw",
		RawQuery: "user=vraj",
		
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("generated url is: ",anotherUrl)
}