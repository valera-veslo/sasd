package main

import (
	"errors"
	"os"
	"reflect"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
}

func (l *LinkedList) Count() int {
	count := 0
	node := l.head
	for node != nil {
		count++
		node = node.next
	}
	return count
}

func (l *LinkedList) Find(n int) (Node, error) {
	node := l.head
	for node != nil {
		if node.value == n {
			return *node, nil
		}
		node = node.next
	}
	return Node{value: -1, next: nil}, errors.New("value not found")
}

func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node
	node := l.head
	for node != nil {
		if node.value == n {
			nodes = append(nodes, *node)
		}
		node = node.next
	}
	return nodes
}

func (l *LinkedList) Delete(n int, all bool) {
	var prev *Node
	node := l.head
	for node != nil {
		if node.value == n {
			if prev == nil {
				l.head = node.next
			} else {
				prev.next = node.next
			}
			if node == l.tail {
				l.tail = prev
			}
			if !all {
				break
			}
		} else {
			prev = node
		}
		node = node.next
	}
}

func (l *LinkedList) Insert(after *Node, add Node) {
	if after == nil {
		return
	}
	if after == l.tail {
		l.AddInTail(add)
		return
	}
	add.next = after.next
	after.next = &add
}

func (l *LinkedList) InsertFirst(first Node) {
	if l.head == nil {
		l.head = &first
		l.tail = &first
	} else {
		first.next = l.head
		l.head = &first
	}
}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
}
