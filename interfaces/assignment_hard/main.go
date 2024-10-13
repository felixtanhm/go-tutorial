package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid file path")
		return
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	res, err := io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	fmt.Println("\n", res)
}
