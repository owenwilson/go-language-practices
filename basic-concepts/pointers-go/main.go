package main

import "fmt"

func main() {
	// pointers: variables that store a memory address
	var color string = "red"
	var pointerColor *string
	pointerColor = &color
	*pointerColor = "blue"
	fmt.Printf("Type: %T, Value: %s, Direction: %v\n", color, color, &color)
	fmt.Printf("Type: %T, Value: %v\n", pointerColor, pointerColor)
	fmt.Printf("Type: %T, Value: %v, Dereferencing: %s\n", pointerColor, pointerColor, *pointerColor)
}
