package ds

import "fmt"

// СТАТТЯ ЯКАСЬ ДИВНА З КУПОЮ ПОМИЛОК, АЛЕ ІНШИХ НЕ ЗНАЙШОВ
// https://dev.to/srleyva/golang-self-balancing-tree-insertions-255b

type RedBlackTree struct {
	Root *RedBlackNode
}

type Color uint

const (
	RED Color = iota
	BLACK
)

type RedBlackNode struct {
	Color  Color
	Key    string
	Parent *RedBlackNode
	Left   *RedBlackNode
	Right  *RedBlackNode
}

var NilNode = &RedBlackNode{Color: BLACK}

func (r *RedBlackTree) Insert(node *RedBlackNode) {
	if r.Root == nil {
		r.Root = node
	} else {
		currentNode := r.Root
		for {
			if node.Key == currentNode.Key {
				return // Do nothing as node already in tree
			}
			if node.Key > currentNode.Key {
				if currentNode.Right == NilNode {
					node.Parent = currentNode
					currentNode.Right = node
					break
				}
				currentNode = currentNode.Right
			}
			if node.Key < currentNode.Key {
				if currentNode.Left == NilNode {
					node.Parent = currentNode
					currentNode.Left = node
					break
				}
				currentNode = currentNode.Left
			}
		}
		r.fixTreeAfterAdd(node) // No-Op right now
	}
}

func (r *RedBlackTree) Search(key string) (*RedBlackNode, error) {
	currentNode := r.Root
	for currentNode != NilNode && key != currentNode.Key {
		if key < currentNode.Key {
			currentNode = currentNode.Left
		}
		if key > currentNode.Key {
			currentNode = currentNode.Right
		}
	}

	if currentNode == NilNode {
		return nil, fmt.Errorf("key does not exist in tree: %s", key)
	}
	return currentNode, nil
}

func (r *RedBlackTree) rotateLeft(node *RedBlackNode) {
	// Right of node (old parent) will become the new parent
	RightNode := node.Right
	// Assign the right tree of new parent to old parent
	node.Right = RightNode.Left
	// Tell the right tree of new parent about the change by setting a new Parent Property
	RightNode.Left.Parent = node
	// Assign old parent's parent to new parent
	RightNode.Parent = node.Parent
	// Correct top of tree

	// If this is the top of the tree
	// Make RightNode root
	if RightNode.Parent == nil {
		r.Root = RightNode
	}
	// If This is not top of tree, need to fix node's parent node
	// Check if node is Left or Right of Parent Node and assign RightNode
	if node == node.Parent.Left {
		node.Parent.Left = RightNode
	}
	if node == node.Parent.Right {
		node.Parent.Right = RightNode
	}
	// Demote current parent to child of new parent
	RightNode.Left = node
	node.Parent = RightNode

}

func (r *RedBlackTree) rotateRight(node *RedBlackNode) {
	LeftNode := node.Left
	node.Left = LeftNode.Right

	LeftNode.Right.Parent = node
	LeftNode.Parent = node.Parent

	if LeftNode.Parent == nil {
		r.Root = LeftNode
	} else if node == node.Parent.Left {
		node.Parent.Left = LeftNode
	} else if node == node.Parent.Right {
		node.Parent.Right = LeftNode
	}

	LeftNode.Right = node
	node.Parent = LeftNode
}

func (r *RedBlackTree) fixTreeAfterAdd(node *RedBlackNode) {
	node.Color = RED // New nodes are always red

	// If node is root or the parent is Black, we don't need to perform any re-coloring
	for node != r.Root && node.Parent.Color == RED {

		// Check if parent is a right or left node

		// If parent is left node
		if node.Parent.Parent.Left == node.Parent {
			// because parent is left of grandparent, uncle is right of grandparent
			uncle := node.Parent.Parent.Right

			// If uncle is red, recolor parent and uncle to black
			if uncle.Color == RED {

				// Recolor parent and uncle to black increasing the black node count from nth node to leaf by 1
				node.Parent.Color = BLACK
				uncle.Color = BLACK

				// Make the grandparent red
				node.Parent.Parent.Color = RED

				// Traverse up the tree to grandparent and re-run through fix algorithm
				node = node.Parent.Parent
			} else { // Uncle is black
				// if node is a right child
				if node == node.Parent.Right {
					// traverse up to parent
					node = node.Parent
					// rotate left
					r.rotateLeft(node)
				}
				// recolor parent to black
				node.Parent.Color = BLACK

				// recolor grandparent to red
				node.Parent.Parent.Color = RED

				// rotate grandparent right
				r.rotateRight(node.Parent.Parent)
			}
		} else if node.Parent.Parent.Right == node.Parent {
			// because parent is right of grandparent uncle is inverse left of grandparent
			uncle := node.Parent.Parent.Left

			// if uncle is read re-coloring is all thats needed
			if uncle.Color == RED {
				// increasing black node count to leaf from nth node to one
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				// traverse to grandparent and re-run fix algorithm
				node = node.Parent.Parent
			} else { // uncle is black, rotation is needed
				if node == node.Parent.Left {
					// traverse to parent
					node = node.Parent

					// rotate right
					r.rotateRight(node)
				}
				// recolor parent to black
				node.Parent.Color = BLACK

				// recolor grandparent to red
				node.Parent.Parent.Color = RED

				// rotate left
				r.rotateLeft(node.Parent.Parent)
			}
		}
	}

	// To satisfy Red Black Tree constraint #1, recolor root to black
	r.Root.Color = BLACK
}
