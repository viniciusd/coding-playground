package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("The must be one and only one argument indicating",
			"the name of the file that will be displayed")
		os.Exit(1)
	}
	filePath := os.Args[1]
	bs, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file: ", filePath)
		os.Exit(1)
	}

	io.Copy(os.Stdout, bytes.NewReader(bs))
}
