package main

import (
	"fmt"
)

func main() {
	// slices: these are pointers to an array, they do not contain data.
	things := [7]string{"car_1", "car_2", "car_3", "car_4", "car_5", "car_6"}
	cars := things[0:4]
	slice_red := things[3:7] // slice_red := things[:]
	slice_red[3] = "car_10"
	fmt.Println("Array cars: ", things)
	fmt.Println("Array cars_slice: ", cars)
	fmt.Println("Array slice_red: ", slice_red)
	fmt.Println("Array slice_red: ", slice_red)
}
