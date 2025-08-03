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

// Return the value of the head of list or nil if list is empty
func (dll *DoublyLinkedList) GetHead() *types.Job {
	if dll.IsEmpty() {
		return nil
	}
	return &dll.head.value
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

// Remove a Job node from the tail of the double linked list and return it
func (dll *DoublyLinkedList) RemoveLast() *Node {

	if dll.IsEmpty() {
		return nil
	}

	// Single node list
	if dll.head == dll.tail {
		headNode := dll.head

		// Clean up links
		headNode.next = nil
		headNode.prev = nil

		// Reset list
		dll.head = nil
		dll.tail = nil
		dll.size = 0

		return headNode
	}

	// Create new tail node
	secondToLastNode := dll.tail.prev
	secondToLastNode.next = nil

	// Clean up links
	tailNode := dll.tail
	tailNode.next = nil
	tailNode.prev = nil

	// Update list
	dll.tail = secondToLastNode
	dll.size -= 1

	return tailNode
}

// Remove a Job node from the head of the double linked list and return it
func (dll *DoublyLinkedList) RemoveFirst() *Node {

	if dll.IsEmpty() {
		return nil
	}

	// Single node list
	if dll.head == dll.tail {
		headNode := dll.head

		// Clean up links
		headNode.next = nil
		headNode.prev = nil

		// Reset list
		dll.head = nil
		dll.tail = nil
		dll.size = 0

		return headNode
	}

	// Create new head node
	secondListNode := dll.head.next
	secondListNode.prev = nil

	// Clean up links
	headNode := dll.head
	headNode.next = nil
	headNode.prev = nil

	// Reset list
	dll.head = secondListNode
	dll.size -= 1

	return headNode
}
