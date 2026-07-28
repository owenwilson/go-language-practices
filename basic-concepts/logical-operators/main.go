package main

import "fmt"

func main() {
	// comparison operators >, <, ==, !=, >=, <=
	fmt.Println("(4+5)<6 ", (4+5) < 6)
	fmt.Println("4==4 ", 4 == 4)
	fmt.Println("4!=4 ", 4 != 4)
	fmt.Println("4>=4 ", 4 >= 4)
	fmt.Println("6>=4 ", 6 >= 4)

	// logical operators && , ||
	var age uint = 72
	fmt.Println("is he an audlt? ", age >= 18 && age <= 60)
	fmt.Println("is he a child or an elderly person? ", age < 18 || age > 70)

	// unary logical operator: !
	fmt.Println("!(4 == 4)", !(4 == 4))
	fmt.Println("!(4 != 4)", !(4 != 4))
}
