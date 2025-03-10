package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{value, nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{value, nil}
	if l.Head == nil {
		l.Head = newNode
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next

	}
	currentNode.Next = newNode

}

func main() {
	list := &LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	fmt.Println(list.Head.Value)
	fmt.Println(list.Head.Next.Value)
	fmt.Println(list.Head.Next.Next.Value)
	fmt.Println(list.Head.Next.Next.Next.Value)
	fmt.Println(list.Head.Next.Next.Next.Next.Value)
}
