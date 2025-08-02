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

// Return true if doubly linked list is empty, false otherwise
func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.head == nil
}

// Return the size of the doubly linked list
func (dll *DoublyLinkedList) GetSize() int {
	return dll.size
}

// Add a new Job node to the end of the linked list
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

	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
	dll.size += 1

	return newNode
}

// Add a new Job node to the beginning of the linked list
func (dll *DoublyLinkedList) Prepend(job types.Job) *Node {

	newNode := &Node{
		value: job,
	}

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
		dll.size = 1
		return newNode
	}

	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode
	dll.size += 1

	return newNode
}
