package main

import "fmt"

type Node struct {
	Value               int
	Left, Right, Parent *Node
}

func NewNode(value int, left, right *Node) *Node {
	n := &Node{Value: value, Left: left, Right: right}
	if left != nil {
		left.Parent = n
	}
	if right != nil {
		right.Parent = n
	}
	return n
}

func NewTerminalNode(value int) *Node {
	return &Node{Value: value}
}

type InOrderIterator struct {
	current, root *Node
	yieldedStart  bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	i := &InOrderIterator{root: root, current: root, yieldedStart: false}
	for i.current.Left != nil {
		i.current = i.current.Left
	}
	return i
}

func (i *InOrderIterator) Reset() {
	i.current = i.root
	i.yieldedStart = false
	for i.current.Left != nil {
		i.current = i.current.Left
	}
}

func (i *InOrderIterator) Next() bool {
	if !i.yieldedStart {
		i.yieldedStart = true
		return true
	}
	if i.current.Right != nil {
		i.current = i.current.Right
		for i.current.Left != nil {
			i.current = i.current.Left
		}
		return true
	} else {
		p := i.current.Parent
		for p != nil && i.current == p.Right {
			i.current = p
			p = p.Parent
		}
		i.current = p
		return i.current != nil
	}
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}

func (b *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(b.root)
}

func main() {
	//   1
	//  / \
	// 2   3

	// in-order: 213
	// pre-order: 123
	// post-order: 231

	root := NewNode(1,
		NewTerminalNode(2),
		NewTerminalNode(3))

	it := NewInOrderIterator(root)
	for it.Next() {
		fmt.Printf("%d ", it.current.Value)
	}
	fmt.Println()

	tree := NewBinaryTree(root)
	it = tree.InOrder()
	for it.Next() {
		fmt.Printf("%d ", it.current.Value)
	}
	fmt.Println()
}

// Tree traversal is a common problem in computer science.
// It is the process of visiting each node in a tree data structure, such as the DOM, and executing a command at each node.
// There are three common ways to traverse a tree: depth-first, breadth-first, and in-order.
// Depth-first traversal visits each node as deep as possible before moving to the next sibling.
// Breadth-first traversal visits each node level by level, moving left to right.
// In-order traversal visits the left branch, then the current node, then the right branch.
