package main

import "fmt"

var print = fmt.Println

func main() {

	print(1)

	var arr = [5]int{1, 2, 3, 4, 5}

	slice1 := make([]string, 6)

	print(arr)
	print(slice1)

	slice := arr[0:2]

	print(slice)

	slice = append(slice, 12)

	print(slice)
	print(arr)

	slice = append(slice, 12)
	slice = append(slice, 12)
	slice = append(slice, 12)
	print(slice)

	slice = append(slice, 12)
	slice = append(slice, 12)
	slice = append(slice, 12)
	print(slice)
	print(arr)

}
