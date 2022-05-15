package main

import "fmt"

func main() {
	var fruitArr [2]string // Array have fixed length but slice can change lengths dynamically
	//assign
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)

	//Declear and assing
	fruitA := [2]string{"Banana", "Nanasi"}
	fmt.Println(fruitA)

	//Slice
	fruitSlice := []string{"Banana", "Nanasi", "Apple", "Mango"}
	fmt.Println(fruitSlice)

	//len(fruit)  --- length
	//range fruitSlice[1:2] or fruitSlice[1:...]

}
