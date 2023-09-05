package main

import "fmt"

var print = fmt.Println

type rectangle struct {
	length float64
	width  float64
}

func area(r rectangle) float64 {
	return r.length * r.width
}

func (r rectangle) areaMethod() float64 {
	return r.width * r.length
}

func main() {

	r := rectangle{12, 10.3}
	print(area(r))

	print(r.areaMethod())
}
