package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	next := bst

	for next != nil {
		if i <= next.data {
			if next.left == nil {
				next.left = NewBst(i)
				return
			}

			next = next.left
		} else {
			if next.right == nil {
				next.right = NewBst(i)
				return
			}

			next = next.right
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	sorted := []int{}

	if bst == nil {
		return sorted
	}

	if bst.left != nil {
		sorted = append(sorted, bst.left.SortedData()...)
	}

	sorted = append(sorted, bst.data)

	if bst.right != nil {
		sorted = append(sorted, bst.right.SortedData()...)
	}

	return sorted
}
