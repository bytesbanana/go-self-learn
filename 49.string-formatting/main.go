package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	fmt.Printf("struct1: %v\n", p)  // General values
	fmt.Printf("struct2: %+v\n", p) // General values + fields name
	fmt.Printf("struct3: %#v\n", p) // General values + fields name + source code snippet

	fmt.Printf("type: %T\n", p)    // Type
	fmt.Printf("bool: %t\n", true) // Boolean
	fmt.Printf("int: %d\n", 123)   // integers

	fmt.Printf("binary: %b\n", 123) // integers
	fmt.Printf("char: %c\n", 33)    // character

	fmt.Printf("hex: %x\n", 33) // hex

	fmt.Printf("float1: %f\n", 78.9)        // floating point
	fmt.Printf("float2: %e\n", 123400000.0) // floating point
	fmt.Printf("float2: %E\n", 123400000.0) // floating point scientific notation.

	fmt.Printf("str1: %s\n", "\"string\"") // Basic string
	fmt.Printf("str2: %q\n", "\"string\"") // Basic string to double quote string
	fmt.Printf("str3: %x\n", "hex this")   // Byte of string

	fmt.Printf("pointer: %p\n", &p) // Pointer

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)         // number width and presision "right-justified"
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)   // floating point width and presision "right-justified"
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // "left-justified"
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")      // string right justified
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")    // string left justified

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") //io.writer

}
