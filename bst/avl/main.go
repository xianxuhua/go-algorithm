package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Node struct {
	Val    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func (t *AVLTree) AddWithTraverse(Val int) {
	t.Root = t.addWithTraverse(t.Root, Val)
}

func (n *Node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

func (n *Node) getBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.getHeight() - n.Right.getHeight()
}

func (n *Node) leftRotate() *Node {
	// node.getBalanceFactor()	   y				x 				  x
	//		   t4      x   ->	y		z     ->  y       z
	//			    t3   z	  t4       t1 t2	t4 t3   t1 t2
	//	     		   t1 t2
	y := n
	x := n.Right
	t3 := x.Left

	x.Left = y
	y.Right = t3

	// 更新height
	y.Height = max(y.Left.getHeight(), y.Right.getHeight()) + 1
	x.Height = max(x.Left.getHeight(), x.Right.getHeight()) + 1
	return n
}

func (n *Node) rightRotate() *Node {
	// 右旋转	   y				x 				  x
	//		   x      t4   ->	z		y     ->  z       y
	//		z	 t3			  t1 t2       t4	t1 t2   t3 t4
	//	  t1 t2
	y := n
	x := y.Left
	t3 := x.Right

	x.Right = y
	y.Left = t3

	// 更新height
	y.Height = max(y.Left.getHeight(), y.Right.getHeight()) + 1
	x.Height = max(x.Left.getHeight(), x.Right.getHeight()) + 1

	return n
}

func (t *AVLTree) addWithTraverse(node *Node, Val int) *Node {
	if node == nil {
		return &Node{Val: Val}
	}

	if node.Val >= Val {
		node.Left = t.addWithTraverse(node.Left, Val)
	} else if node.Val <= Val {
		node.Right = t.addWithTraverse(node.Right, Val)
	}
	node.Height = max(node.Left.getHeight(), node.Right.getHeight()) + 1

	// 维护平衡
	// LL
	if node.getBalanceFactor() > 1 && node.Left.getBalanceFactor() >= 0 {
		return node.rightRotate()
	}

	// RR
	if node.getBalanceFactor() < -1 && node.Right.getBalanceFactor() <= 0 {
		return node.leftRotate()
	}

	// LR
	if node.getBalanceFactor() > 1 && node.Left.getBalanceFactor() < 0 {
		node.Left = node.Left.leftRotate()
		return node.rightRotate()
	}

	// RL
	if node.getBalanceFactor() < -1 && node.Right.getBalanceFactor() > 0 {
		node.Right = node.Right.rightRotate()
		return node.leftRotate()
	}

	return node
}

func (t *AVLTree) MidTraverse(opt func(int)) {
	t.midTraverse(t.Root, opt)
}

func (t *AVLTree) midTraverse(now *Node, opt func(int)) {
	if now == nil {
		return
	}

	t.midTraverse(now.Left, opt)
	opt(now.Val)
	t.midTraverse(now.Right, opt)
}

func (t *AVLTree) IsBST() bool {
	Vals := []int{}
	t.MidTraverse(func(i int) {
		Vals = append(Vals, i)
	})
	for i := 0; i < len(Vals)-1; i++ {
		if Vals[i+1] < Vals[i] {
			return false
		}
	}

	return true
}

func (t *AVLTree) IsAVL() bool {
	return t.isAVL(t.Root)
}

func (t *AVLTree) isAVL(node *Node) bool {
	if node == nil {
		return true
	}

	balanceFactor := node.getBalanceFactor()
	if math.Abs(float64(balanceFactor)) > 1 {
		return false
	}

	return t.isAVL(node.Left) && t.isAVL(node.Right)
}

func main() {
	t := AVLTree{}
	for i := 0; i < 1000; i++ {
		t.AddWithTraverse(rand.Intn(1000))
	}
	fmt.Println(t.IsBST())
	fmt.Println(t.IsAVL())
}
