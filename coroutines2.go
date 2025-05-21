package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func wait(s string) {

	for i := 0; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		pl(len(s))
	}

}

func main() {

	go wait("aaaaaaa")

	wait("qwert")
}
