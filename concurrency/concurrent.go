package main

import (
	"fmt"
	"os"
)

// calcNums calculates hexadecimal numbers in the range [start, end]
// and writes them to the provided file f. Once done, it signals the
// completion on the doneCh channel.
func calcNums(start, end int, f *os.File, doneCh chan struct{}) {
	for i := start; i <= end; i++ {
		_, err := f.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}
	}
	doneCh <- struct{}{}
}

func main() {
	// Create a file named "file.txt" for writing
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	numGoRoutines := 10
	doneCh := make(chan struct{}) // Channel to signal completion of calcNums goroutines

	final := 16777215
	for i := 0; i <= final; i += (final / numGoRoutines) + 1 {
		paso := i + (final / numGoRoutines)

		if paso > final {
			paso = final
		}
		fmt.Printf("ejecutando %d %d\n", i, paso)

		// Start a goroutine to calculate hexadecimal numbers
		go calcNums(i, paso, f, doneCh)
	}

	doneNum := 0
	for doneNum < numGoRoutines {
		<-doneCh // Wait for a calcNums goroutine to complete
		fmt.Printf("terminó uno\n")
		doneNum++
	}

	fmt.Println("Listo!") // Print "Listo!" indicating the program has finished
}
