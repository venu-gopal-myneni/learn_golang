package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// 0, 0.0, false - defaults- can't change length

	var arr1 [5]int

	arr1[0] = 34

	arr2 := [5]int{345, 35, 123, 1232, 123213}

	pl(arr2[1])
	pl(len(arr2))

	pl(arr1)
	pl(arr2)

	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	for i, v := range arr2 {
		pl(i, v)
	}

	// multi dime arrays

	arr3 := [2][3]int{
		{1, 2, 3},
		{5, 66, 657},
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			pl(arr3[i][j])
		}
	}
}
