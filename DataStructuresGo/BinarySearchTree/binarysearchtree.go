package BinarySearchTree

type Node struct {
	Data       int
	leftChild  *Node
	rightChild *Node
}

type BinarySearchTree struct {
	Root  *Node
	count int
}

func NewTree() *BinarySearchTree {
	tree := BinarySearchTree{Root: nil, count: 0}
	return &tree
}

func (root *Node) Search(value int) *Node {
	if root == nil {
		return nil
	}
	ptrNode := root
	if ptrNode.Data == value {
		return ptrNode
	} else if value < ptrNode.Data {
		return ptrNode.leftChild.Search(value)
	} else if value > ptrNode.Data {
		return ptrNode.rightChild.Search(value)
	}
	return nil
}

func (tree *BinarySearchTree) Insert(value int) {

	if tree.Root == nil {
		tree.Root = &Node{Data: value, leftChild: nil, rightChild: nil}
	} else {
		ptrNode := tree.Root
		if ptrNode.Data == value {

		}
	}

}