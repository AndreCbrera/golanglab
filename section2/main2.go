package main

import "fmt"

// arrays
func todo() {
	// var arr []int
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}
	arr2 = append(arr2, "is", "Andres", "!")
	fmt.Println(arr, arr2)
}

func main() {
	todo()
}