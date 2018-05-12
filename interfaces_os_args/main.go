package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("The must be one and only one argument indicating",
			"the name of the file that will be displayed")
		os.Exit(1)
	}
	filePath := os.Args[1]
	f, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error reading file: ", filePath)
		os.Exit(2)
	}

	io.Copy(os.Stdout, f)
}
