package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func MakeSlice(n int) []string {
	out := []string{}
	for i := 0; i < n; i++ {
		out = append(out, strconv.Itoa(i))
	}
	return out
}

func main() {
	var intSlice []int
	var strSlice []string
	strSlice = make([]string, 0)

	slice := MakeSlice(20)

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(slice))
	fmt.Println(reflect.ValueOf(strSlice).Kind(), len(strSlice))

}
