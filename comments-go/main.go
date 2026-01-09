// It is the main package
// where our applications runs
package main

import "fmt"

// example const
const domain = "mydomain.net"

// example variable
var name = "owen-wilson"

// example large comment
const (
	Jan = iota + 1
	Feb
	Mar
	Abr
	May
	Jun
)

func main() {
	// function print
	fmt.Println(domain, name)
}
