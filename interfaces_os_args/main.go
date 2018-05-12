package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, input := range os.Args[1:] {
		if input == "-" {
			io.Copy(os.Stdout, os.Stdin)
			continue
		}

		f, err := os.Open(input)
		if err != nil {
			fmt.Println("Error reading file: ", input)
			os.Exit(1)
		}
		io.Copy(os.Stdout, f)
	}
}
