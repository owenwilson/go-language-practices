package main

import "fmt"

func main() {
	// plugin vim-go
	fmt.Println("vim-go")
	// example 1 variables
	var cargo_truck_1 string
	cargo_truck_1 = "nissan"
	var cargo_truck_2 string
	cargo_truck_2 = "volvo"

	// example 2 variables
	var laptop_brand_1 string = "toshiba"
	var laptop_brand_2 string = "dell"

	// example 3 variables
	var (
		apple  string = "manzana"
		banana string = "banana"
		orange string = "naranja"
	)

	// example 4 variables
	var city_bo_1, city_bo_2 string = "la paz", "oruro"

	// example 5 variable
	phone_brand_1, phone_brand_2 := "honor", "xiaomi"
	// variables of one data type cannot be changed to another data type, for exapmle, from string to integer
	phone_brand_1, phone_brand_3 := "realme", "iphone"

	// pring variables
	fmt.Println(cargo_truck_1, cargo_truck_2)
	fmt.Println(laptop_brand_1, laptop_brand_2)
	fmt.Println(city_bo_1, city_bo_2)
	fmt.Println(apple, banana, orange)
	fmt.Println(phone_brand_1, phone_brand_2, phone_brand_3)
}
