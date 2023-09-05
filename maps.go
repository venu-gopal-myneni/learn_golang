package main

import "fmt"

var print = fmt.Println

func main() {
	var heroes map[string]string
	heroes = make(map[string]string)

	villians := make(map[string]string)

	heroes["batman"] = "Brucy"
	villians["penguin"] = "Prbgy"

	superPets := map[int]string{1: "fdsaf", 345: "werewr"}

	print(heroes)
	print(villians)
	print(superPets)
}
