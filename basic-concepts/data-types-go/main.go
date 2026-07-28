package main

import "fmt"

func main() {
	// bool, string, numeric
	var a bool = true

	fmt.Printf("\nType: %T, Value: %v\n", a, a)

	// string
	var nameString string = "my-example"
	fmt.Printf("\nType: %T, Value: %v\n", nameString, nameString)

	// numeric
	var age uint8 = 33
	fmt.Printf("\nType: %T, Value: %v\n", age, age)

	// byte
	var data byte = 12
	fmt.Printf("\nType: %T, Value: %v\n", data, data)

	// rune, example unicode a = 97
	var my_var rune = 'a'
	fmt.Printf("\nType: %T, Value: %v\n", my_var, my_var)

	// float32
	var my_var_float float32 = 123.24
	fmt.Printf("\nType: %T, Value: %v\n", my_var_float, my_var_float)

	// remember that you can't add a uint to an int
	// for example
	var add_var_a uint8 = 255
	var add_var_b uint16 = 2550

	fmt.Println("\n#### identifier blank ####")
	//c := uint16(add_var_a) + add_var_b
	_ = uint16(add_var_a) + add_var_b
	//fmt.Printf("Type: %T, Value: %v\n", c, c)
	fmt.Printf("Type: %T, Value: %v\n", add_var_a, add_var_a)
	fmt.Println("#### end identifier blank ####")

	// example value empty
	var var_example string
	fmt.Println("\n#### value empty ####")
	fmt.Printf("Type: %T. Value: %v\n", var_example, var_example)
	fmt.Printf("Type: %T. Value: %q\n", var_example, var_example)
	fmt.Println("#### end value empty ####")

	// example value 0
	var var_example_uint uint8

	fmt.Println("\n#### value 0 ####")
	fmt.Printf("Type: %T, Value: %v\n", var_example_uint, var_example_uint)
	fmt.Println("#### end value 0 ####")

	// example default value bool is false
	var var_example_boolean bool
	fmt.Println("\n#### value bool ####")
	fmt.Printf("Type: %T, Value: %v\n", var_example_boolean, var_example_boolean)
	fmt.Printf("Type: %T, Value: %q\n", var_example_boolean, var_example_boolean)
	fmt.Println("#### end value bool ####")
}
