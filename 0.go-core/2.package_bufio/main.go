package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Example: Reading a file line by line
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Reading lines from example.txt:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading file:", err)
	}

	// Example: Writing to a file using bufio.Writer
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	writer.WriteString("This is an example of bufio.Writer.\n")
	writer.WriteString("It helps with buffered writes to a file.\n")
	writer.Flush()

	fmt.Println("Data successfully written to output.txt")
}
