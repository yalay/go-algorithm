package main

import (
	"fmt"
)

type MyList struct {
	head *MyNode
}

type MyNode struct {
	value int
	next  *MyNode
}

func (m *MyList) String() {
	if m == nil {
		return
	}

	fmt.Printf("%d ", m.head.value)
	nextNode := m.head.next
	for {
		if nextNode == nil {
			fmt.Println("")
			return
		}

		fmt.Printf("%d ", nextNode.value)
		nextNode = nextNode.next
	}
}

// 反转链表
func (m *MyList) Reverse() {
	p1 := m.head
	p2 := p1.next
	p1.next = nil
	for {
		if p2 == nil {
			m.head = p1
			return
		}
		p3 := p2.next
		p2.next = p1

		p1 = p2
		p2 = p3
	}
	m.head = p1
}

// 递归创建一个数值从value到max的链表
func newMyNode(value, max int) *MyNode {
	if value > max {
		return nil
	}

	return &MyNode{
		value: value,
		next:  newMyNode(value+1, max),
	}
}

func main() {
	newList := &MyList{
		head: newMyNode(1, 10),
	}
	newList.String()

	newList.Reverse()
	newList.String()
}
