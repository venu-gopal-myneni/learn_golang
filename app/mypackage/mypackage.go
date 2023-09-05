package mypackage

import (
	"strconv"
)

var MyName string = "Venu"

func IntArrToStrArr(intArr []int) []string {
	var strArray []string

	for _, inti := range intArr {
		strArray = append(strArray, strconv.Itoa(inti))
	}

	return strArray
}
