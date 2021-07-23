package main

import (
	"fmt"

	"github.com/ghadad/my-go-training/utils"
)

func Sum(arr []int, size int, sum int) int {
	if size == 0 {
		return sum
	}
	//fmt.Println(arr, size, sum)
	return Sum(arr[:size], size-1, arr[size-1]+sum)
}

func main() {
	s := utils.RandomIntSlice(20)
	fmt.Println("sum of slice ", s, ":", Sum(s, len(s), 0))
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("sum of slice ", s2, ":", Sum(s2, len(s2), 0))
}
