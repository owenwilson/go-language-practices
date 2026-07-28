package main

import "fmt"

func main() {
	// arithmetic operators (), *, /, %, +, -
	var example_add = (2 + 3) * 2
	fmt.Println("result: ", example_add)

	// assignment operators =, +=, -=, *=, /=, %=
	var var_b int = 5
	// version 1
	//var_b = var_b + 2
	//fmt.Println("result: ", var_b)
	// version 2
	var_b += 2
	fmt.Println("result: ", var_b)

	// post-increment and post-decrement statements
	// are not expressions but statements
	var var_c int = 6
	var_c++
	fmt.Println("result: ", var_c)
	var_c--
	fmt.Println("result: ", var_c)
}
