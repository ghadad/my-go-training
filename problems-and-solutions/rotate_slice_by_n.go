package main

import (
	"fmt"

	"github.com/ghadad/my-go-training/utils"
)

func RotateLeftByN(s []int, places int) []int {
	res := s
	for i := 0; i < places; i++ {
		res = append(res, res[i])
	}
	return res[places:len(res)]
}

func RotateRightByN(s []int, places int) []int {
	res := s
	slice_len := len(s)
	for i := 0; i < places; i++ {
		res[0] = res[slice_len-1]
	}
	return res[places:len(res)]
}

func main() {

	s := utils.RandomIntSlice(10)
	places := 5
	fmt.Println("Rotate left ", s, " by ", places, "=>", RotateLeftByN(s, places))
	places = 1
	fmt.Println("Rotate  left ", s, " by ", places, "=>", RotateLeftByN(s, places))
}
