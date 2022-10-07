package main

import (
	"fmt"
	"os"
)

func main() {

	file := createFile("/tmp/defer.txt")
	defer closeFile(file) // Ensure this function call is perform after program executed
	// Usaully use for cleanup
	// NOTE: It is a "finally" in other langauge

	writeFile(file)
}

func createFile(path string) *os.File {
	fmt.Println("creating a new file")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f

}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintf(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v \n", err)
		os.Exit(1)
	}
}
