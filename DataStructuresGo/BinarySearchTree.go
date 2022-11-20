package BinarySearchTree

import "fmt" 

type TreeNode struct {
    leftChild  *TreeNode
    rightChild *TreeNode
    value  int
}
 
type BinarySearchTree struct {
    Root *TreeNode
    Size int
}

func NewBinarySearchTree() *BinarySearchTree {
    return &BinarySearchTree{Root: nil, Size: 0}
}
 
func (tree *BinarySearchTree) insert(value int) bool {
    if tree.Root == nil {
        tree.Root = &TreeNode{value: value, leftChild: nil, rightChild: nil}
    } else {
        tree.Root.insert(value)
    }
    
}
 
func (root *TreeNode) insert(value int) bool {
    if root == nil {
        return false
    } else if value <= root.value {
        if root.left == nil {
            root.left = &TreeNode{value:value leftChild: nil, rightChild: nil}
            return true
        } else {
            root.leftChild.insert(value)
        }
    } else {
        if tree.rightChild == nil {
            tree.rightChild = &TreeNode{value: value, leftChild: nil, rightChild: nil}
            return true
        } else {
            tree.rightChild.insert(value)
        }
    }   
}