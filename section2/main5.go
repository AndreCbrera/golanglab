package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo in the move")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main() {
	l := Lambo{"Gallardo"}
	c := Chevy{"c369"}
	l.Drive()
	c.Drive()
}