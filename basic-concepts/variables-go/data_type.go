package main

import "fmt"

func main() {
	// bool, string, numeric
	var word_1 string = "golang language"

	// uint8 - limit 0-255
	var a uint8 = 255

	// uint16 - limit 0-65535
	var b uint16 = 65535

	// uint32 - limit 0-4294967295
	var c uint32 = 4294967295

	// uint64 - limit 0-18446744073709551615
	var d uint64 = 18446744073709551615

	// byte is an alias for uint8 - limit 0-255
	var byte_a byte = 255

	// rune is an alias for uint32
	var rune_a rune = 'a'
	var rune_b rune = 1550

	// float
	var float_variable = 132.24
	var balance float32 = 123.4

	fmt.Printf("Type: %T, Value: %v\n", word_1, word_1)
	fmt.Printf("Type: %T, Value: %v\n", a, a)
	fmt.Printf("Type: %T, Value: %v\n", b, b)
	fmt.Printf("Type: %T, Value: %v\n", c, c)
	fmt.Printf("Type: %T, Value: %v\n", d, d)
	fmt.Printf("Type: %T, Value: %v\n", byte_a, byte_a)
	fmt.Printf("Type: %T, Value: %v\n", rune_a, rune_a)
	fmt.Printf("Type: %T, Value: %v\n", rune_b, rune_b)
	fmt.Printf("Type: %T, Value: %v\n", float_variable, float_variable)
	fmt.Printf("Type: %T, Value: %v\n", balance, balance)

	// example two sum numbers
	var value_a uint8 = 255
	var value_b uint16 = 2550
	result := uint16(value_a) + value_b

	// operaton blank identifier
	_ = uint16(value_a) + value_b

	fmt.Println("\n#### SUM ####")
	fmt.Printf("Type: %T, Value: %v\n", result, result)
	fmt.Printf("Type: %T, Value: %v\n", value_a, value_b)

	// empty chain
	var chain string

	fmt.Printf("Type: %T, Value: %q\n", chain, chain)

	// empty integer
	var number_integer uint8

	fmt.Printf("Type: %T, Value: %v\n", number_integer, number_integer)

	// empty boolean
	var value_bool bool
	fmt.Printf("Type: %T, Value: %v\n", value_bool, value_bool)
}
