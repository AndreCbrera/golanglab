package main

import "fmt"

/* declare global variable */
var g int = 20

func main()  {
	/* Declare local variable */
	var g int = 10

	/* The local variable within the function are priorized */

	fmt.Printf("Value: g = %d\n", g)
} 