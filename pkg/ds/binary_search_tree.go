package ds

// Read this https://appliedgo.net/bintree/

// Another implementation

type node struct {
	left  *node
	data  int
	right *node
}
type bst struct {
	root *node
}

type BSTService interface {
	Insert(data int) *bst
	GetRoot() *node
	Search(key int) bool
}

func (tree *bst) Insert(data int) *bst {
	if tree.root == nil {
		tree.root = &node{
			data: data,
		}
	} else {
		tree.root.insert(data)
	}
	return tree
}

func (nd *node) insert(data int) {
	if nd == nil {
		return
	} else if data < nd.data {
		if nd.left == nil {
			nd.left = &node{
				data: data,
			}
		} else {
			nd.left.insert(data)
		}
	} else {
		if nd.right == nil {
			nd.right = &node{
				data: data,
			}
		} else {
			nd.right.insert(data)
		}
	}
}

func (tree *bst) Search(key int) bool {
	var keyFound bool
	if tree.root == nil {
		keyFound = false
	} else if tree.root.data != key {
		keyFound = tree.root.search(key)
	} else {
		keyFound = true
	}
	return keyFound
}
func (nd *node) search(key int) bool {
	var keyFound bool
	if nd == nil {
		keyFound = false
	} else if key == nd.data {
		keyFound = true
	} else {
		if nd.data > key {
			keyFound = nd.left.search(key)
		} else {
			keyFound = nd.right.search(key)
		}
	}
	return keyFound
}
func (tree *bst) GetRoot() *node {
	return tree.root
}
