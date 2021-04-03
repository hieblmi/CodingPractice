package main

import (
	"fmt"
)

var root Node

func FindTwoSum(target int) bool {
	root = Node{Val: 5}
	root.Left = &Node{Val: 3}
	root.Left.Left = &Node{Val: 2}
	root.Left.Right = &Node{Val: 4}
	root.Right = &Node{Val: 7}
	root.Right.Right = &Node{Val: 8}
	root.Right.Left = &Node{Val: 6}

	chLeft := make(chan int)
	chRight := make(chan int)

	go searchLeftBranch(&root, chLeft)
	go searchRightBranch(&root, chRight)

	var leftVal, rightVal, sum int
	leftVal = <-chLeft
	rightVal = <-chRight
	sum = leftVal + rightVal
	for sum != target {
		if sum < target {
			leftVal = <-chLeft
			sum = leftVal + rightVal
		} else {
			rightVal = <-chRight
			sum = leftVal + rightVal
		}
		if leftVal == rightVal {
			return false
		}
	}
	res := sum == target
	if res {
		fmt.Printf("%d + %d = %d\n", leftVal, rightVal, sum)
	}
	return res
}

func searchRightBranch(n *Node, ch chan int) {
	if n == nil {
		return
	}

	searchRightBranch(n.Right, ch)
	ch <- n.Val
	searchRightBranch(n.Left, ch)
}

func searchLeftBranch(n *Node, ch chan int) {
	if n == nil {
		return
	}

	searchLeftBranch(n.Left, ch)
	ch <- n.Val
	searchLeftBranch(n.Right, ch)
}
