package main

import (
	"os"

	"github.com/HamzaMasood1/learning-go/datastructures/trees/binaryTree"
)

func main() {
	tree := &binaryTree.BinaryTree{}
	tree.Insert(100).
		Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)
	print(os.Stdout, tree.Root, 0, 'M')
}
