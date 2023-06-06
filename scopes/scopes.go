package main 

import "fmt"

var x = 7 // global scope

func exmaple() {
	fmt.Println(x)
}

func main()  {
	//x := 7 // local scope
	fmt.Println(x)
	exmaple() // this example will show 7 7
}