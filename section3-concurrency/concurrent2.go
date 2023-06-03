package main

import (
	"fmt"
	"os"
	"strings"
)

// calcNums calculates hexadecimal numbers in the range [start, end]
// and sends the generated numbers as a string to the resultCh channel.
// Once done, it signals the completion on the doneCh channel.
func calcNums(start, end int, resultCh chan string, doneCh chan struct{}) {
	var sBuilder strings.Builder
	for i := start; i <= end; i++ {
		fmt.Fprint(&sBuilder, fmt.Sprintf("%06x\n", i))
	}
	resultCh <- sBuilder.String()
	doneCh <- struct{}{}
}

func main() {
	// Create a file named "file.txt" for writing
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	// Create a channel to receive strings containing hexadecimal numbers
	outCh := make(chan string)

	// Create a channel to signal completion of the writer goroutine
	doneWrite := make(chan struct{})

	// Writer goroutine
	go func() {
		for s := range outCh {
			// Write the generated numbers to the file
			_, err := f.WriteString(s)
			if err != nil {
				panic(err)
			}
		}
		doneWrite <- struct{}{} // Signal completion of the writer goroutine
	}()

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
		go calcNums(i, paso, outCh, doneCh)
	}

	doneNum := 0
	for doneNum < numGoRoutines {
		<-doneCh // Wait for a calcNums goroutine to complete
		fmt.Printf("terminó uno\n")
		doneNum++
	}

	close(outCh)  // Close the outCh channel to signal completion to the writer goroutine
	<-doneWrite   // Wait for the writer goroutine to complete
	fmt.Println("Listo!") // Print "Listo!" indicating the program has finished
}
