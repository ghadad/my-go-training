package main

import (
	"fmt"
	"math"
)

func multi10(result int, pow int) int {
	if pow == 1 {
		return result
	}
	return multi10(result*10, pow-1)
}

func ReverseNumber(n int) int {
	stack := make([]int, 0)
	for n > 0 {
		rem := n % 10
		n -= rem
		n = n / 10
		stack = append(stack, rem)
	}

	sum := 0

	stack_len := len(stack)

	for i := 0; i < stack_len; i++ {

		sum += stack[i] * int(math.Pow10((stack_len-i)-1))
	}
	return sum
}

func main() {
	fmt.Println(1367, " .. reverse =>", ReverseNumber(1367))
}
