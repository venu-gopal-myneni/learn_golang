package main

import (
	mypackage "examples/package/mypackage"
	"fmt"
)

func main() {

	fmt.Println(123, mypackage.MyName)

	intArr := []int{1, 2, 3, 4, 5, 66}

	fmt.Println(mypackage.IntArrToStrArr(intArr))
}
