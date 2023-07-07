package binaryTree

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	Root *BinaryNode
}

func (t *BinaryTree) Insert(data int64) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.Root.Insert(data)
	}
	return t
}

func (n *BinaryNode) Insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.Insert(data)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}
