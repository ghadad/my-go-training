package main

import "fmt"

type Stack struct {
	stack []int
	count int
}

func (s *Stack) Push(k int) {
	s.stack = append(s.stack, k)
	s.count++
}

func (s *Stack) Pop() interface{} {
	stackLen := len(s.stack)
	if stackLen == 0 {
		return nil
	}
	element := s.stack[stackLen-1]
	s.stack = s.stack[:stackLen-1]
	return element
}

func main() {
	s := &Stack{}
	i := 0
	for i < 10 {
		s.Push(i)
		i++
	}
	fmt.Println(s)

	for i < 25 {
		elm := s.Pop()
		fmt.Println("Pop element:", elm)
		i++
	}

	fmt.Println(s)

}
