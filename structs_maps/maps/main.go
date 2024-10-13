package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	printColors(colors)
}

func printColors(colors map[string]string) {
	for color, hex := range colors {
		fmt.Printf("Hexcode for %v is %v\n", color, hex)
	}
}
