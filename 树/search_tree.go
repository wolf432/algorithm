package main

import (
	"fmt"
)

type SearchTree struct {
	data int
	left *SearchTree
	right *SearchTree
}

func newSearchTree(data int) *SearchTree {
	return &SearchTree{
		data: data,
		left: nil,
		right: nil,
	}
}

func (root *SearchTree) insert(data int){
	p := root
	for p != nil{
		if data > p.data{
			if p.right == nil {
				p.right = newSearchTree(data)
				break
			}
			p = p.right
		}else{
			if p.left == nil{
				p.left = newSearchTree(data)
				return
			}
			p = p.left
		}
	}
}

func (root *SearchTree) delete(data int) bool{
	p := root          //指向要删除的节点，初始化指向根节点
	var pp *SearchTree //记录的是p的父节点

	for p != nil && p.data != data{
		pp = p
		if data > p.data{
			p = p.right
		}else{
			p = p.left
		}
	}
	if p == nil{ //没找到
		return false
	}

	//要删除的节点有两个子节点
	if p.left != nil && p.right != nil{ //查找右子树中最小节点
		minP := p.right
		minPP := p
		for minP.left != nil{
			minPP = minP
			minP = minP.left
		}
		p.data = minP.data
		p = minP
		pp = minPP
	}

	//删除节点是叶子节点或者仅有一个子节点
	var child *SearchTree
	if p.left != nil{
		child = p.left
	}else if p.right != nil {
		child = p.right
	}else{
		child = nil
	}

	if pp == nil{ //要删除的是根节点
		root = child
	}else if pp.left == p{
		pp.left = child
	}else{
		pp.right = child
	}
	return true
}

func (root *SearchTree) find(data int) (*SearchTree, bool){
	p := root
	for p != nil{
		if data < p.data{
			p = p.left
		}else if data > p.data {
			p = p.right
		}else{
			return p, true
		}
	}
	return nil, false
}


//前序遍历:即直接按照我们对结点的访问顺序输出遍历结果即实现，父结点值被最先输出
func (root *SearchTree) preOrder() {
	fmt.Println(root.data)
	if root.left != nil{
		root.left.preOrder()
	}
	if root.right != nil{
		root.right.preOrder()
	}
}

//中序遍历:左孩子值最先输出，然后是父结点，最后是右孩子
func (root *SearchTree) inOrder(){
	if root.left != nil{
		root.left.preOrder()
	}
	fmt.Println(root.data)
	if root.right != nil{
		root.right.preOrder()
	}
}

//后序遍历：左右孩子值依次输出，最后是父结点
func (root *SearchTree) postOrder(){
	if root.left != nil{
		root.left.preOrder()
	}
	if root.right != nil{
		root.right.preOrder()
	}
	fmt.Println(root.data)
}


func main() {
	root := newSearchTree(4)

	root.left = newSearchTree(2)
	root.right = newSearchTree(7)

	root.left.left = newSearchTree(1)
	root.left.right = newSearchTree(3)
	root.right.left = newSearchTree(6)
	root.right.right = newSearchTree(9)

	sort(root)
}

func sort(root *SearchTree){
	if root == nil{
		return
	}

	sort(root.left)
	fmt.Println(root.data)
	sort(root.right)
}