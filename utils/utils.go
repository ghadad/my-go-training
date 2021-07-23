package utils

import (
	"math/rand"
	"time"
)

func RandomIntSlice(size int) []int {
	rand.Seed(int64(time.Now().Nanosecond()))

	slice := make([]int, 0)
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(1000))

	}
	return slice
}

func RandomCharSlice(size int) []byte {

	slice := make([]byte, 0)
	var c byte
	for i := 0; i < size; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		c = 'A'
		slice = append(slice, c+byte(rand.Intn(26)))

	}
	return slice
}
