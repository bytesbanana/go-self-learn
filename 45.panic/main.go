package main

import "os"

func main() {

	panic("a problem")

	// This part will be abort
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
