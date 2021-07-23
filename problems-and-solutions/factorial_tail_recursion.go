package main

import "fmt"

func FactorialTr(i int, result int) int {
	if i == 0 {
		return result
	}

	return FactorialTr(i-1, i*result)
}
func main() {
	fmt.Println("Factorial of ", 5, " = ", FactorialTr(5, 1))
	fmt.Println("Factorial of ", 15, " = ", FactorialTr(15, 1))
}
