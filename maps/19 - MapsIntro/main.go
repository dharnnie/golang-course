package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"black": "#000000",
	// }
	// var colors map[string]string // 2nd way
	colors := make(map[string]string)
	colors["red"] = "#FF0000" // append a new value to an existing map
	fmt.Println(colors)
	delete(colors, "red")
	fmt.Println(colors)
}
