package main

import (
	"github.com/qingsong-he/ce"
	"github.com/qingsong-he/some/golang/t_bitree/bitree"
)

func printPreOrder(node *bitree.BiTreeNode) {
	if node != nil {
		println("printPreOrder", node.Data().(int))
		if node.Left() != nil {
			printPreOrder(node.Left())
		}
		if node.Right() != nil {
			printPreOrder(node.Right())
		}
	}
}

func printInOrder(node *bitree.BiTreeNode) {
	if node != nil {
		if node.Left() != nil {
			printInOrder(node.Left())
		}
		println("printInOrder", node.Data().(int))
		if node.Right() != nil {
			printInOrder(node.Right())
		}
	}
}

func printPostOrder(node *bitree.BiTreeNode) {
	if node != nil {
		if node.Left() != nil {
			printPostOrder(node.Left())
		}
		if node.Right() != nil {
			printPostOrder(node.Right())
		}
		println("printPostOrder", node.Data().(int))
	}
}

func searchInt(biTree *bitree.BiTree, ival int) *bitree.BiTreeNode {
	var node = biTree.Root()
	for {
		if node == nil {
			break
		}

		meta := node.Data().(int)

		if ival == meta {
			return node
		} else if ival < meta {
			node = node.Left()
		} else {
			node = node.Right()
		}
	}
	return nil
}

func insertInt(biTree *bitree.BiTree, ival int) {
	var direction int
	var pre *bitree.BiTreeNode
	var node = biTree.Root()
	for {
		if node == nil {
			break
		}

		pre = node

		meta := node.Data().(int)
		if ival == meta {
			return
		} else if ival < meta {
			node = node.Left()
			direction = 1
		} else {
			node = node.Right()
			direction = 2
		}
	}

	switch direction {
	case 0:
		biTree.InsertLeft(nil, ival)
	case 1:
		biTree.InsertLeft(pre, ival)
	case 2:
		biTree.InsertRight(pre, ival)
	}
}

func main() {
	b1 := bitree.New()
	insertInt(b1, 20)
	insertInt(b1, 10)
	insertInt(b1, 30)
	insertInt(b1, 15)
	insertInt(b1, 25)
	insertInt(b1, 70)
	insertInt(b1, 80)
	insertInt(b1, 23)
	insertInt(b1, 26)
	insertInt(b1, 5)
	ce.Print(b1.Size())

	n := searchInt(b1, 20)
	if n != nil {
		ce.Print(n.Data())
	}
	printPreOrder(b1.Root())
	printInOrder(b1.Root())
	printPostOrder(b1.Root())
}
