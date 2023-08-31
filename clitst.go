package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	args := os.Args[1:]

	for _, val := range args {
		fmt.Println(val)
		ind, err := strconv.Atoi(val)

		if err != nil {
			panic(err)
		}

		fmt.Println(ind)
	}

}
