package utils

import (
	"http-job-que-system/types"
)

type Node struct {
	value types.Job
	next  *Node // Can be nil (no next node)
	prev  *Node // Can be nil (no previous node)
}

type DoublyLinkedList struct {
	head *Node // Points to first node, or nil if empty
	tail *Node // Points to last node, or nil if empty
	size int
}

func CreateNewDoubleLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Returns true if doubly linked list is empty, false otherwise
func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.head == nil
}

func (dll *DoublyLinkedList) Append(job types.Job) *Node {
	newNode := &Node{
		value: job,
	}

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
		dll.size = 1
		return newNode
	}

	dll.tail.next = newNode
	newNode.prev = dll.tail
	dll.tail = newNode
	dll.size += 1

	return newNode
}

// func (dll *DoublyLinkedList) Prepend(job types.Job) *Node {
// 	// Implementation
// }
//
// func (dll *DoublyLinkedList) GetSize() int {
// 	// Implementation
// }
