package main

import (
	"fmt"
)

// structure encapsulation
type Car struct {
	Name	string
	Age		int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("Driving...")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	c := Car{
		Name: 	 "Chevy",
		Age: 	 20,
		ModelNo: 20132,
	}
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}