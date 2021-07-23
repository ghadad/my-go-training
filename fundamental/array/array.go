package main

import "fmt"

const ArraySize = 10

func Initilze() [ArraySize]int {
	var arr [ArraySize]int
	return arr
}

var Arr2 [3]int = [...]int{1, 2, 3}

func GetArr2() [3]int {
	return Arr2
}
func main() {
	fmt.Println(Initilze())
}
