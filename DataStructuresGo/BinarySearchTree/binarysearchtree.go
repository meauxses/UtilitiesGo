package BinarySearchTree

type TreeNode struct {
	Data       int
	leftChild  *TreeNode
	rightChild *TreeNode
}

type BinarySearchTree struct {
	Root  *TreeNode
	count int
}

func NewTree() *BinarySearchTree {
	tree := BinarySearchTree{Root: nil, count: 0}
	return &tree
}

func (root *TreeNode) Search(value int) *TreeNode {
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
//** need to complete */
func (tree *BinarySearchTree) Insert(value int) {

	if tree.Root == nil {
		tree.Root = &TreeNode{Data: value, leftChild: nil, rightChild: nil}
	} else {
		ptrNode := tree.Root
		if ptrNode.Data == value {

		}
	}

}