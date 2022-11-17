package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":    "stop",
		"yellow": "wait",
		"green":  "go",
	}
	//fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, value := range m {
		fmt.Println(color, value)
	}
}
