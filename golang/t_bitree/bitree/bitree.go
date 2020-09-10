package bitree

import (
	"errors"
)

type BiTreeNode struct {
	data  interface{}
	left  *BiTreeNode
	right *BiTreeNode
}
type BiTree struct {
	size    int
	root    *BiTreeNode
	compare func(a, b interface{}) int
}

func NewBiTree() *BiTree {
	return &BiTree{
		size: 0,
		root: nil,
	}
}

var ErrByLogical = errors.New("logical error")

func (b *BiTree) InsertLeft(n *BiTreeNode, data interface{}) {
	var pos **BiTreeNode

	if n == nil {
		if b.size > 0 {
			panic(ErrByLogical)
		}
		pos = &(b.root)
	} else {
		if n.left != nil {
			panic(ErrByLogical)
		}
		pos = &(n.left)
	}

	*pos = &BiTreeNode{
		data: data,
	}
	b.size++
}

func (b *BiTree) InsertRight(n *BiTreeNode, data interface{}) {
	var pos **BiTreeNode

	if n == nil {
		if b.size > 0 {
			panic(ErrByLogical)
		}
		pos = &(b.root)
	} else {
		if n.right != nil {
			panic(ErrByLogical)
		}
		pos = &(n.right)
	}

	*pos = &BiTreeNode{
		data: data,
	}
	b.size++
}

func (b *BiTree) RemoveLeft(n *BiTreeNode) {
	if b.size == 0 {
		return
	}

	var pos **BiTreeNode
	if n == nil {
		pos = &(b.root)
	} else {
		pos = &(n.left)
	}
	if *pos != nil {
		b.RemoveLeft(*pos)
		b.RemoveRight(*pos)

		(*pos).data = nil
		*pos = nil

		b.size--
	}
}

func (b *BiTree) RemoveRight(n *BiTreeNode) {
	if b.size == 0 {
		return
	}

	var pos **BiTreeNode
	if n == nil {
		pos = &(b.root)
	} else {
		pos = &(n.right)
	}
	if *pos != nil {
		b.RemoveLeft(*pos)
		b.RemoveRight(*pos)

		(*pos).data = nil
		*pos = nil

		b.size--
	}
}

func (b *BiTree) Root() *BiTreeNode {
	return b.root
}

func (b *BiTree) Size() int {
	return b.size
}

func (b *BiTreeNode) Data() interface{} {
	return b.data
}

func (b *BiTreeNode) Left() *BiTreeNode {
	return b.left
}

func (b *BiTreeNode) Right() *BiTreeNode {
	return b.right
}

func BiTreeMerge(left, right *BiTree, data interface{}) *BiTree {
	n := NewBiTree()
	n.InsertLeft(nil, data)

	if left == nil {
		panic(ErrByLogical)
	}
	if right == nil {
		panic(ErrByLogical)
	}

	n.root.left = left.root
	n.root.right = right.root
	n.size = n.size + left.size + right.size

	return n
}
