package main

import "fmt"

// defing a function with return value
func returnIntegerValue() (int, int) {
	return 1, 2
}

func main() {
	a, _ := returnIntegerValue()
	fmt.Println(a)
}
