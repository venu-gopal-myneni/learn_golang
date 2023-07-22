package main

import "fmt"

var print = fmt.Println

func sayHello() {
	print("Hello")
}

func getOne(x int, y int) int {
	return x + y
}

func getTwo(x int) (int, int) {
	return x, x * x
}

func getQuot(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("Y is zero")
	} else {
		return x / y, nil
	}
}

// varidaic functions
func getSum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func arraySum(arr []int) int {

	sum := 0

	for _, num := range arr {
		sum += num
	}

	return sum

}

// pass by pointers
func changeval(ptr *int) {
	*ptr = 200
}
func main() {
	// func funcName(params) returnType {BODY}
	sayHello()
	print(getOne(1, 2))
	print(getTwo(3))

	print(getQuot(100, 2))

	print(getSum(4, 4234, 23432423, 34324))

	arr := []int{1, 2, 3, 3, 2, 1}

	print("array sum :", arraySum(arr))

	ptr := 20
	print("before:", ptr)
	changeval(&ptr)
	print("after: ", ptr)
}
