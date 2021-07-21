package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	key   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(k int) *Tree {
	if t.root == nil {
		t.root = &Node{key: k}
		return t
	} else {
		t.root.Insert(k)
	}
	return t
}

func (n *Node) Insert(k int) {

	if n.key < k {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}
	} else if n.key > k {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

func (t *Tree) Traverse() {
	if t.root == nil {
		fmt.Println("Tree is empty")
	}
	t.root.Traverse()
}

func (n *Node) Traverse() {
	fmt.Println(n.key)
	if n.right != nil {
		n.right.Traverse()
	}
	if n.left != nil {
		n.left.Traverse()
	}

}

func (t *Tree) Search(k int) bool {
	if t.root == nil {
		fmt.Println("Tree is empty")
		return false
	}
	return t.root.Search(k)
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.key > k {
		return n.left.Search(k)
	}
	if n.key < k {
		return n.right.Search(k)
	}
	return true
}

func Filter(m map[int]bool, b bool) map[int]bool {
	result := make(map[int]bool)

	for k, v := range m {
		if v == b {
			result[k] = v
		}
	}
	return result
}

func main() {
	t := new(Tree)
	maxRand := 50

	for i := 0; i < maxRand; i++ {
		t.Insert(rand.Intn(maxRand))
	}

	//	t.Traverse()
	rand.Seed(time.Now().UnixNano())

	result := make(map[int]bool)
	for i := 0; i < maxRand; i++ {
		n := rand.Intn(maxRand)
		result[n] = t.Search(n)
	}

	fmt.Println("Found elements :", Filter(result, true))
	fmt.Println("Not found elements :", Filter(result, false))
}
