package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	// if else, for, switch case, break continue

	fmt.Println("-----------")

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	fmt.Println("-----------")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------")

	mymap := make(map[string]interface{})
	mymap["name"] = "Andres"
	mymap["age"] = 20

	for k, v :=  range mymap {
		fmt.Printf("Key: %s and value: %v", k, v)
	}
}