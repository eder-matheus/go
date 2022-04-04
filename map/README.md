package main

import "fmt"

func main() {
	// first map declaration syntax
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// second map declaration
	// var colors map[string]string

	// third map declaration
	colors := make(map[string]string)

	colors["white"] = "#ffffff"

	// remove the entry on the key informed on the second parameter
	delete(colors, "white")

	fmt.Println(colors)
}
