package main

import (
	"fmt"

	"github.com/ghadad/my-go-training/utils"
)

func Reverse(arr []int) []int {
	i, j := 0, len(arr)-1
	for ; i < j; i++ {
		arr[i], arr[j] = arr[j], arr[i]
		j--
	}

	return arr
}

func main() {

	s := utils.RandomIntSlice(10)
	fmt.Println(s)
	fmt.Println("=======> Reverse : =======>\n", Reverse(s))

}
