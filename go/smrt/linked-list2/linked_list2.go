package main

import (
	"errors"
	"os"
	"reflect"
)

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail
	}

	l.tail = &item
	l.tail.next = nil
}

func (l *LinkedList2) Count() int {
	count := 0
	node := l.head
	for node != nil {
		count++
		node = node.next
	}
	return count
}

func (l *LinkedList2) Find(n int) (Node, error) {
	node := l.head
	for node != nil {
		if node.value == n {
			return *node, nil
		}
		node = node.next
	}
	return Node{value: -1, next: nil}, errors.New("value not found")
}

func (l *LinkedList2) FindAll(n int) []Node {
	var nodes []Node
	lnode := l.head
	rnode := l.tail
	if lnode == nil {
		return nodes
	}
	if lnode == rnode {
		if lnode.value == n {
			nodes = append(nodes, *lnode)
		}
		return nodes
	}
	for lnode != rnode {
		if lnode.value == n {
			nodes = append(nodes, *lnode)
		}
		if rnode.value == n {
			nodes = append(nodes, *rnode)
		}
		if lnode.next == nil {
			break
		}
		if lnode.next == rnode {
			break
		}
		if lnode.next == rnode.prev {
			if lnode.next.value == n {
				nodes = append(nodes, *lnode.next)
			}
			break
		}
		lnode = lnode.next
		rnode = rnode.prev
	}
	return nodes
}

func (l *LinkedList2) Delete(n int, all bool) {
	node := l.head
	nodes := make([]*Node, 0)
	for node != nil {
		if node.value == n {
			nodes = append(nodes, node)
			if !all {
				break
			}
		}
		node = node.next
	}
	for i := range nodes {
		node = nodes[i]
		if node.next != nil {
			node.next.prev = node.prev
		} else {
			l.tail = node.prev
		}
		if node.prev != nil {
			node.prev.next = node.next
		} else {
			l.head = node.next
		}
	}
}

func (l *LinkedList2) Insert(after *Node, add Node) {
	if after == nil {
		return
	}
	if after == l.tail {
		l.AddInTail(add)
		return
	}
	add.next = after.next
	add.prev = after
	after.next.prev = &add
	after.next = &add
}

func (l *LinkedList2) InsertFirst(first Node) {
	if l.head == nil {
		l.head = &first
		l.tail = &first
	} else {
		first.next = l.head
		l.head.prev = &first
		l.head = &first
	}
}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
}
