package main

import (
	"fmt"
	"time"
)

var print = fmt.Println

func print1to15() {
	for i := 1; i <= 15; i++ {
		print("func 1", i)
	}
}

func print1to10() {
	for i := 1; i <= 10; i++ {
		print("func 2", i)
	}
}

func main() {
	print(1 + 2)
	go print1to10()
	go print1to15()
	time.Sleep(2 * time.Second)
}
