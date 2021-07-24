package main

import (
	"fmt"

	"github.com/ghadad/my-go-training/utils"
)

func FindBalancePoint(s []int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	leftSum := 0
	rightSum := sum

	for i := 0; leftSum <= rightSum; i++ {
		fmt.Println("position:", i, "compare sum :", leftSum, "< == > ", rightSum)
		rightSum -= s[i]
		if leftSum == rightSum {
			fmt.Println("compare sum :", leftSum, "< == > ", rightSum)
			return i
		}
		leftSum += s[i]
	}
	return 0
}

func main() {
	s := utils.RandomIntSlice(15)
	fmt.Println(s, " => balace position :", FindBalancePoint(s))
	s = []int{5, 5, 5, 1, 5, 5, 5}
	fmt.Println(s, FindBalancePoint(s))

}
