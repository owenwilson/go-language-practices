package main

import "fmt"

func main() {
	// example arrays
	var flags [3]string
	flags[0] = "flag_1"
	flags[1] = "flag_2"
	flags[2] = "flag_3"
	//flags[3] = "flag_4"
	fmt.Println("Arrays flags: ", flags)

	// example arrays 2
	example_flags := [3]string{"e_flag_1", "e_flag_2", "e_flag_3"}
	fmt.Println("Arrays_example_2:", example_flags)

	// example arrays 3
	accessories := [...]string{"mouse", "phone", "table"}
	fmt.Println("Arrays_example_3: ", accessories)
}
