package main

import "fmt"

func main() {
	// string variable
	var a = "initial"
	fmt.Println("#### STRING #####\n", a)
	fmt.Printf("Type var a: %T\n", a)
	fmt.Println("#################\n")

	// integer variable
	var b, c = 1, 2
	fmt.Println("#### INTEGER ####\n", b, c)
	fmt.Printf("Type var b: %T\n", b)
	fmt.Printf("Type var c: %T\n", c)
	fmt.Println("#################\n")

	// boolean variable
	var d = true
	fmt.Println("#### BOOLEAN ####\n", d)
	fmt.Printf("Type: var d: %T\n", d)
	fmt.Println("#################\n")

	// variable declared without a corresponding initialization are zero-valued
	// For example, the zero value for an int is 0
	var e int
	fmt.Println("#### THE ZERO VALUE FOR AN INT IS 0 ####\n", e)
	fmt.Println("Type var e: %T\n", e)
	fmt.Println("########################################")

	// the following syntax is only
	// available inside functions
	f := "apple"
	fmt.Println("#### := ####\n", f)
	fmt.Println("Type var f: ", f)
	fmt.Println("############")
}
