package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {

	second := l.head
	l.head = n
	l.head.next = second
	l.length++

}

func (l LinkedList) printList() {

	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

func main() {
	l := LinkedList{}
	node1 := &node{data: 5}
	node2 := &node{data: 7}
	node3 := &node{data: 9}

	l.prepend(node1)
	l.prepend(node2)
	l.prepend(node3)

	fmt.Println(l)
	l.printList()

}

