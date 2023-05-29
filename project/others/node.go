package main

import (
	"fmt"
	"log"
	"math"
)

// Node 链表
type Node struct {
	val  int
	next *Node
}

func newNode(val int) *Node {
	return &Node{val: val}
}
func newNodeNil() *Node {
	return nil
}

func main() {
	testIntersectNode()
}

func testIsLoop() {
	var node = newNode(1)
	node.next = newNode(2)
	node.next.next = newNode(3)
	node.next.next.next = newNode(4)
	node.next.next.next.next = node.next

	n := isLoop(node)
	if n == nil {
		log.Fatalln("no loop")
	}

	fmt.Println(n.val)
}

func testIntersectNode() {
	var node1 = newNode(8)
	node1.next = newNode(11)
	node1.next.next = newNode(3)
	node1.next.next.next = newNode(4)
	node1.next.next.next.next = newNode(5)
	node1.next.next.next.next.next = newNode(77)
	node1.next.next.next.next.next.next = node1.next.next.next

	var node2 = newNode(1)
	node2.next = newNode(2)
	node2.next.next = node1.next
	node2.next.next.next = newNode(7)
	node2.next.next.next.next = newNode(8)
	node2.next.next.next.next.next = newNode(9)
	node2.next.next.next.next.next.next = node2.next.next.next.next

	n := getIntersectNode(node1, node2)
	if n == nil {
		log.Fatalln("no loop")
	}

	fmt.Println(n.val)
}

func getIntersectNode(node1, node2 *Node) *Node {
	// 给定两个点，判断是否相交
	loop1 := isLoop(node1)
	loop2 := isLoop(node2)

	if loop1 == nil && loop2 == nil {
		return noLoop(node1, node2)
	}

	if loop1 != nil && loop2 != nil {
		return bothLoop(node1, loop1, node2, loop2)
	}

	return nil
}

func noLoop(node1, node2 *Node) *Node {
	var cur1, cur2, n = node1, node2, 0

	for cur1.next != nil {
		n++
		cur1 = cur1.next
	}

	for cur2.next != nil {
		n--
		cur2 = cur2.next
	}

	if cur1 != cur2 {
		return nil
	}

	if n > 0 { //谁长，谁的头变成cur1
		cur1 = node1
	} else {
		cur1 = node2
	}
	// 谁短，谁的头变成cur2
	if cur1 == node1 {
		cur2 = node2
	} else {
		cur2 = node1
	}

	n = int(math.Abs(float64(n)))
	for n != 0 {
		n--
		cur1 = cur1.next
	}

	for cur1 != cur2 {
		cur1 = cur1.next
		cur2 = cur2.next
	}

	return cur1
}

func bothLoop(node1, loop1, node2, loop2 *Node) *Node {
	var cur1, cur2, n = newNodeNil(), newNodeNil(), 0
	if loop1 == loop2 {
		cur1 = node1
		cur2 = node2
		for cur1.next != loop1 {
			n++
			cur1 = cur1.next
		}
		for cur2.next != loop2 {
			n--
			cur2 = cur2.next
		}
		if n > 0 { //谁长，谁的头变成cur1
			cur1 = node1
		} else {
			cur1 = node2
		}
		// 谁短，谁的头变成cur2
		if cur1 == node1 {
			cur2 = node2
		} else {
			cur2 = node1
		}
		n = int(math.Abs(float64(n)))
		for n != 0 {
			n--
			cur1 = cur1.next
		}

		for cur1 != cur2 {
			cur1 = cur1.next
			cur2 = cur2.next
		}
		return cur1
	} else {
		cur1 = loop1.next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.next
		}
		return nil
	}
}

func isLoop(node *Node) *Node {
	//是否成环
	if node == nil || node.next == nil || node.next.next == nil {
		return nil
	}
	var s, f = node.next, node.next.next
	for s != f {
		s = s.next
		f = f.next.next
		if f.next == nil || f.next.next == nil {
			return nil
		}
	}

	f = node
	for s != f {
		s = s.next
		f = f.next
	}

	return s
}
