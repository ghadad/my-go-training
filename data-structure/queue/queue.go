package main

import "fmt"

type Queue struct {
	head  []int
	count int
}

func (q *Queue) Enqueue(i int) {
	q.head = append(q.head, i)
	q.count++
}

func (q *Queue) Dequeue() interface{} {
	if len(q.head) == 0 {
		return nil
	}

	v := q.head[0]
	fmt.Println(v)
	q.head = q.head[1:]
	return v
}

func (q *Queue) Print() {

}
func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Printf("%v\n", q)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	a := q.Dequeue()
	fmt.Printf("%v %v\n", q, a)
}
