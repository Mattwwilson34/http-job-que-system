# Project Notes

## What is a Que?
A data structure that follow a first in first out principle (FIFO).

ex. [a,b,c] => + d => [a,b,c,d] => pop que => [b,c,d]

What data needs to be stored on the que data struct?

Queue = {
    items = stores the items of our que
    enqueue = add new item at tail of the queue
    dequeue = remove item at the head of the queue
    peek = return the item at the head of the queue without removing it
}

The items container can be a simple array or a doubly linked list

## Double Linked List
DoublyLinkedList = {
    next = job_2
    prev = null
    value = job_1

    insertBefore = insert a new node before an existing node in the list
    insertAfter = insert a new node after an existing node in the list
    insertBeginning = insert a new node into an empty list
    insertEnd = insert a new node at the end of the list
    remove = remove an existing node from the list
}

DoublyLinkedList = {
    firstNode = job_1
    lastNode = job_5
}
