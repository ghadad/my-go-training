package main

import (
	"fmt"

	"github.com/ghadad/my-go-training/utils"
)

func UniqeAndCount(s []byte) map[byte]int {
	var m = make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	return m
}

func main() {
	s := utils.RandomCharSlice(1000)
	fmt.Printf("%s\n", s)
	fmt.Println(UniqeAndCount(s))
}
