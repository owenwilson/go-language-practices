package main

import "fmt"

const (
	fileRegular   = iota //0
	fileDirectory        //1
	fileSymlink          //2
)

func main() {
	fmt.Println("fileRegular: ", fileRegular)     // 0
	fmt.Println("fileDirectory: ", fileDirectory) // 1
	fmt.Println("fileSymlink: ", fileSymlink)     // 2
}
