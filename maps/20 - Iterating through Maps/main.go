package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"black": "#000000",
		"white": "#000000",
	}
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("%s : %s ", key, value)
	}
}
