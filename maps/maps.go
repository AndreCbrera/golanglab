package main

import "fmt"

func main()  {

	ages := make(map[string]int)
	ages["Andres"] = 28
	fmt.Printf("Andres is %d years old\n", ages["Andres"])

	ages["Nayade"] += 1
	fmt.Printf("Nayade is %d years old\n", ages["Nayade"])

	// addr := &ages["Nayade"]

	gpas := map[string]float32 {
		"Andres": 3.4,
		"Nayade": 3.9,
	}

	fmt.Printf("Andres' GPA is %.2f\n", gpas["Andres"])
	fmt.Printf("Nayade' GPA is %.2f\n", gpas["Nayade"])

	var visited map[string]bool
	visited = make(map[string]bool)
	visited["A"] = true
	fmt.Printf("A has been visited? %v\n", visited["A"])

	fruits := []string {
		"bananas",
		"kiwis",
		"apples",
		"tomatoes",
		"strawberries",
		"bananas",
	}
	fmt.Printf("Duplicate fruits? %s\n", findDuplicates(fruits))
}

func findDuplicates(words []string) string {
	dupMap := make(map[string]bool)

	for i := 0; i < len(words); i++{
		if !dupMap[words[i]] {
			dupMap[words[i]] = true
		} else {
			return words[i]
		}
	}
	return("")
}