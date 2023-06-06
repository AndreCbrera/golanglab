package main

import "fmt"

/* Declare global variable */
var a int = 20

func main()  {
	/* main() local variables */
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main a = %d\n", a);
	c = sum(a,b);
	fmt.Printf("main c = %d\n", c);
}

func sum(a, b int) int {
	fmt.Printf("sum() int the function a = %d\n", a)
	fmt.Printf("sum() int the function b = %d\n", b)
	return a + b
}