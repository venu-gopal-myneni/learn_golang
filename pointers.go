package main

import "fmt"

var pl = fmt.Println

func zeroval(name int) {
	name = 0
}

func zeroptr(name *int) {
	*name = 0
}
func main() {
	pl(123)

	i := 20

	pl(i)

	zeroval(i)
	pl(i)

	zeroptr(&i)
	pl(i)

	pl(&i)

}
