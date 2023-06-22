package ds

import "fmt"

type NodeSLL struct {
	name string
	next *NodeSLL
}

type singleList struct {
	len  int
	head *NodeSLL
}

func (s *singleList) AddFront(node *NodeSLL) {
	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
	s.len++
	return
}

func (s *singleList) AddBack(node *NodeSLL) {
	if s.head == nil {
		s.head = node
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	s.len++
	return
}

func (s *singleList) RemoveFront() error {
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}

func (s *singleList) RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}
	var prev *NodeSLL
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}
	s.len--
	return nil
}

func (s *singleList) Front() (string, error) {
	if s.head == nil {
		return "", fmt.Errorf("single List is empty")
	}
	return s.head.name, nil
}

func (s *singleList) Size() int {
	return s.len
}

func (s *singleList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("tranverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}
