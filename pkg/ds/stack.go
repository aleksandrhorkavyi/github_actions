package ds

import (
	"container/list"
	"fmt"
	"sync"
)

//List Implementation

type customStack struct {
	stack *list.List
}

func (c *customStack) Push(value string) {
	c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("pop Error: Stack is empty")
}

func (c *customStack) Front() (string, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peep Error: Stack Datatype is incorrect")
	}
	return "", fmt.Errorf("peep Error: Stack is empty")
}

func (c *customStack) Size() int {
	return c.stack.Len()
}

func (c *customStack) Empty() bool {
	return c.stack.Len() == 0
}

// Slice Implementation
// 1
func foo() {
	var stack []string

	stack = append(stack, "world!") // Push
	stack = append(stack, "Hello ")

	for len(stack) > 0 {
		n := len(stack) - 1 // Top element
		fmt.Print(stack[n])

		stack = stack[:n] // Pop
	}
}

// 2
type ItemType interface{}

// Stack - Stack of items.
type Stack struct {
	// Slice of type ItemType, it holds items in stack.
	items []ItemType
	// rwLock for handling concurrent operations on the stack.
	rwLock sync.RWMutex
}

// Push - Adds an Item to the top of the stack
func (stack *Stack) Push(t ItemType) {
	//Initialize items slice if not initialized
	if stack.items == nil {
		stack.items = []ItemType{}
	}
	// Acquire read, write lock before inserting a new item in the stack.
	stack.rwLock.Lock()
	// Performs append operation.
	stack.items = append(stack.items, t)
	// This will release read, write lock
	stack.rwLock.Unlock()
}

// Pop removes an Item from the top of the stack
func (stack *Stack) Pop() *ItemType {
	// Checking if stack is empty before performing pop operation
	if len(stack.items) == 0 {
		return nil
	}
	// Acquire read, write lock as items are going to modify.
	stack.rwLock.Lock()
	// Popping item from items slice.
	item := stack.items[len(stack.items)-1]
	//Adjusting the item's length accordingly
	stack.items = stack.items[0 : len(stack.items)-1]
	// Release read write lock.
	stack.rwLock.Unlock()
	// Return last popped item
	return &item
}

// Size return size i.e. number of items present in stack.
func (stack *Stack) Size() int {
	// Acquire read lock
	stack.rwLock.RLock()
	// defer operation of unlock.
	defer stack.rwLock.RUnlock()
	// Return length of items slice.
	return len(stack.items)
}

// All - return all items present in stack
func (stack *Stack) All() []ItemType {
	// Acquire read lock
	stack.rwLock.RLock()
	// defer operation of unlock.
	defer stack.rwLock.RUnlock()
	// Return items slice to caller.
	return stack.items
}

// IsEmpty - Check is stack is empty or not.
func (stack *Stack) IsEmpty() bool {
	// Acquire read lock
	stack.rwLock.RLock()
	// defer operation of unlock.
	defer stack.rwLock.RUnlock()
	return len(stack.items) == 0
}

// 3

type customStack2 struct {
	stack []string
	lock  sync.RWMutex
}

func (c *customStack2) Push(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, name)
}

func (c *customStack2) Pop() error {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:len-1]
		return nil
	}
	return fmt.Errorf("pop Error: Stack is empty")
}

func (c *customStack2) Front() (string, error) {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[len-1], nil
	}
	return "", fmt.Errorf("peep Error: Stack is empty")
}

func (c *customStack2) Size() int {
	return len(c.stack)
}

func (c *customStack2) Empty() bool {
	return len(c.stack) == 0
}
