package main

import "fmt"

var print = fmt.Println

type MyConstraint interface {
	int | float64
}

func add(a int, b int) int {
	return a + b
}

func add2[MC MyConstraint](a MC, b MC) MC {
	return a + b
}

func main() {

	print(add(1, 2))
	print(add2(3, 6.77))

}
