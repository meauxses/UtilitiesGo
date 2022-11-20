package LinkedList

// This is a doubly linked list with head and tail nodes. It includes methods for both queue/deque and push/pop for flexibility.

import (
	"fmt"
)

// This is the node object with next and prev pointers
type Node[T comparable] struct {
	Data T
	next *Node[T]
	prev *Node[T]
}

//This linked list has a reference to the head and the tail 
type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	length int
}

//This is a constructor for a new list. Head and Tail are set to nil. It returns a pointer to the list
func NewList[T comparable]() *LinkedList[T] {
	list := &LinkedList[T]{Head: nil, Tail: nil, length: 0}
	return list
}

//This adds a given value to the beginning of the list. This is synonymous with 'queueing' in a queue.
func (l *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Data: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.next = l.Head
		l.Head.prev = newNode
		l.Head = newNode
	}
	l.length++
}

//This adds a given value to the end of the list. This is synonymous with 'pushing' in a stack.
func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Data: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.next = newNode
		newNode.prev = l.Tail
		l.Tail = newNode
	}
	l.length++
}

//This is an internal method for finding the node and index for a given value. This function is used for remove methods. 
func (l *LinkedList[T]) find(value T) (nodeToReturn *Node[T], indexToReturn int) {
	nodeToReturn = nil
	indexToReturn = 0
	found := false
	finder := l.Head
	for finder != nil {
		if finder.Data == value {
			nodeToReturn = finder
			found = true
			break
		} else {
			finder = finder.next
			indexToReturn++
		}
	}
	if !found {
		indexToReturn = -1
	}
	return nodeToReturn, indexToReturn
}

//This is for finding a value at a given index
func (l *LinkedList[T]) FindAtIndex(index int) T {
	if index >= l.length || index < 0 {
		fmt.Println("Index is out of bounds. Returned zeroValue of list type.")
		var zeroValue T 
		return zeroValue
	}
	finder := l.Head
	for i := 0; i < index; i++ {
		finder = finder.next
	}
	return finder.Data
}

//TODO: FindIndexOf(givenvalue)

//This removes the entry with the given value from the list and returns the index where found
func (l *LinkedList[T]) Remove(value T) int {
	foundNode, foundIndex := l.find(value)
	if l.length <= 0 {
		fmt.Println("List is empty.")
		return -1
	}
	if foundIndex >= 0 {
		if foundNode == l.Head && foundNode == l.Tail {
			l.Head = nil
			l.Tail = nil
			foundNode = nil
		} else if foundNode == l.Head {
			foundNode = foundNode.next
			foundNode.prev = nil
			l.Head = foundNode
		} else if foundNode == l.Tail {
			foundNode = foundNode.prev
			foundNode.next = nil
			l.Tail = foundNode
		} else {
			foundNode.prev.next = foundNode.next
			foundNode.next.prev = foundNode.prev
			foundNode = nil
		}
		l.length--
	} else {
		fmt.Println("Value is not contained in list.")
	}
	return foundIndex
}

//This 'pops' an entry from the end of the list. It returns the value removed.
func (l *LinkedList[T]) Pop() T {
	if l.IsEmpty() {
		fmt.Println("List is empty. Returned zeroValue of list type.")
		var zeroValue T
		return zeroValue
	}
	valueToReturn := l.Tail.Data
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil

	} else {
		newTail := l.Tail.prev
		newTail.next = nil
		l.Tail = newTail
	}
	l.length--
	return valueToReturn
}

//This deques an entry from the beginning of the list. It returns the value removed. 
func (l *LinkedList[T]) Dequeue() T {
	if l.IsEmpty() {
		fmt.Println("List is empty. Returned zeroValue of list type.")
		var zeroValue T
		return zeroValue
	}
	valueToReturn := l.Head.Data
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		newHead := l.Head.next
		newHead.prev = nil
		l.Head = newHead
	}
	l.length--
	return valueToReturn
}

//This returns the size of the list
func (l *LinkedList[T]) Size() int {
	return l.length
}

//This returns true if the list is empty
func (l *LinkedList[T]) IsEmpty() bool {
	return l.length == 0
}

//This prints the contents of the list with '|' separators
func (l *LinkedList[T]) Print() {
	if l.IsEmpty() {
		fmt.Println("List is empty.")
	} else {
		printer := l.Head
		fmt.Print("|")
		for printer != nil {
			fmt.Print(printer.Data, "|")
			printer = printer.next
		}
		fmt.Println("")
	}
}
