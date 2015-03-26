package tree

import (
	"errors"
)

type BST struct {
	root *Node
}

// Since Go has no built-in support for Generics,
// here we make it simple to support only int key type
func New() *BST {
	bst := new(BST)
	bst.root = nil
	return bst
}

func (bst *BST) IsEmpty() bool {
	if bst.root == nil {
		return true
	}
	return bst.root.numChildren == 0
}

func (bst *BST) Size() int {
	if bst.root == nil {
		return 0
	}
	return bst.root.numChildren
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.numChildren
}

func (bst *BST) Contains(k Key) bool {
	return bst.Get(k) != nil
}

func (bst *BST) Get(k Key) interface{} {
	return get(bst.root, k)
}

func get(node *Node, key Key) interface{} {
	if node == nil {
		return nil
	}

	if key == node.key {
		return node.value
	} else if key > node.key {
		return get(node.right, key)
	} else {
		return get(node.left, key)
	}
}

func (bst *BST) Put(key Key, value interface{}) {
	if value == nil {
		bst.Delete(key)
		return
	}
	bst.root = put(bst.root, key, value)
}

/*
 Insert the key/value, if the key already exists, update its value
*/
func put(node *Node, key Key, value interface{}) *Node {
	if node == nil {
		return MakeNode(key, value, 1)
	}

	if key > node.key {
		node.right = put(node.right, key, value)
	} else if key < node.key {
		node.left = put(node.left, key, value)
	} else {
		node.value = value
	}
	node.numChildren = 1 + size(node.left) + size(node.right)

	return node
}

func (bst *BST) Delete(key Key) {
	bst.root = delete(bst.root, key)
}

/**
 * Deletion of a node on BST can be accomplished with below 4 steps:
 * 1. Save a link to the node to be deleted;
 * 2. Set x to point to its successor min(tmp.right);
 * 3. Set the right link of x (which is supposed to point to the BST
 *   containing all the keys larger than x.key) to deleteMin(tmp.right),
 *   the link to the BST containing all the keys that are larger than x.key
 *   after the deletion;
 * 4. Set the left link of x (which was nil) to tmp.left (all the keys)
 *  that are less than both the deleted key and its successor.
 */
func delete(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = delete(node.left, key)
	} else if key > node.key {
		node.right = delete(node.right, key)
	} else {
		if node.right == nil {
			return node.left
		}
		if node.left == nil {
			return node.right
		}

		tmp := node
		node = min(tmp.right)
		node.right = deleteMin(tmp.right)
		node.left = tmp.left
	}
	node.numChildren = size(node.left) + size(node.right) + 1
	return node
}

func deleteMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}

	node.left = deleteMin(node.left)
	node.numChildren = size(node.left) + size(node.right) + 1
	return node
}

func (bst *BST) DeleteMin() error {
	if bst.IsEmpty() {
		return errors.New("Cannot delete from an empty BST")
	}

	bst.root = deleteMin(bst.root)
	return nil
}

func min(node *Node) *Node {
	if node.left == nil {
		return node
	} else {
		return min(node.left)
	}
}

func (bst *BST) Min() interface{} {
	if bst.IsEmpty() {
		return nil
	}
	return min(bst.root).key
}

func (bst *BST) Max() interface{} {
	if bst.IsEmpty() {
		return nil
	}
	return max(bst.root).key
}

func max(node *Node) *Node {
	if node.right == nil {
		return node
	} else {
		return max(node.right)
	}
}

func (bst *BST) DeleteMax() error {
	if bst.IsEmpty() {
		return errors.New("Cannot delete element from an empty BST")
	}

	bst.root = deleteMax(bst.root)
	return nil
}

func deleteMax(node *Node) *Node {
	if node.right == nil {
		return node.left
	}
	node.right = deleteMax(node.right)
	node.numChildren = size(node.left) + size(node.right) + 1
	return node
}

func (bst *BST) IsBST() bool {
	return isBST(bst.root, nil, nil)
}

func isBST(node *Node, min, max interface{}) bool {
	if node == nil {
		return true
	}

	if min != nil && node.key < min.(Key) {
		return false
	}
	if max != nil && node.key > max.(Key) {
		return false
	}
	return isBST(node.left, min, node.key) && isBST(node.right, node.key, max)
}
