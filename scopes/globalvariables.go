package main 

import "fmt"

var g int

func main()  {
	
	/* define local variables */
	var a, b int

	a = 10
	b = 20
	g = a + b

	fmt.Printf("Variables: a = %d, b = %d and g = %d\n", a, b, g)
}