package main

import (
	"fmt"
	"sort"
)

func FindMissing(s []int) []int {
	sort.Ints(s)
	i := 0
	result := make([]int, 0)
	for i < len(s)-1 {
		prev := s[i]
		next := s[i+1]

		for j := prev + 1; j < next; j++ {
			//fmt.Println(j, " is missing")
			result = append(result, j)
		}
		i++
	}
	return result
}
func main() {
	slice := []int{13, 1, 2, 3, 7, 8, 9, 0, 14}
	fmt.Println("Given slice :", slice)
	fmt.Println("missing elements :", FindMissing(slice))
}
