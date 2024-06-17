package main

import "fmt"

type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

func newNode(key int) *Node {
	return &Node{
		key:    key,
		height: 1,
		left:   nil,
		right:  nil,
	}
}

func insertNode(n *Node, key int) *Node {
	if n == nil {
		n = newNode(key)
		return n
	}
	search := true
	c := n
	v := make([]*Node, 0)
	for search {
		switch {
		case key < c.key:
			v = append(v, c)
			if c.left == nil {
				c.left = newNode(key)
				search = false
			} else {
				c = c.left
			}
		case key > c.key:
			v = append(v, c)
			if c.right == nil {
				c.right = newNode(key)
				search = false
			} else {
				c = c.right
			}
		}
		for i := len(v) - 1; i >= 0; i-- {
			node := v[i]
			//  fmt.Println(node.height)
			//	fmt.Println(node, ": ",height(node.left), "-", height(node.right))
			//	fmt.Println(newHeight(node))
			node.height = newHeight(node)
			balance := height(node.left) - height(node.right)
			switch {
			case balance > 1:
				fmt.Println("Looking at: ", node)
				if node.left.right == nil {
					rotateRight(node)
				} else if node.left.left == nil {
					rotateLeft(node.left)
					rotateRight(node)
				}

			case balance < -1:
				if node.right.left == nil {
					rotateLeft(node)
				} else if node.left.right == nil {
					rotateRight(node.right)
					rotateLeft(node)
				}
			}
		}
	}
	return n
}

func rotateRight(n *Node) {
	t := n.left
	t1 := n.left.right
	t.right = n
	n.left = t1
	n.height = newHeight(n)
	t.height = newHeight(t)
	fmt.Println("*******")
	fmt.Println(t)
	fmt.Println(n)
	fmt.Println("*******")
}

func rotateLeft(n *Node) {
	t := n.right
	t1 := n.right.left
	t.left = n
	n.right = t1
	n.height = newHeight(n)
	t.height = newHeight(t)
}

func newHeight(n *Node) int {
	return max(height(n.left), height(n.right)) + 1
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printTree(root *Node, indent string, last bool) {
	if root != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += "   "
		} else {
			fmt.Print("L----")
			indent += "|  "
		}
		fmt.Println(root.key, "HEIGHT: ", root.height)
		printTree(root.left, indent, false)
		printTree(root.right, indent, true)
	}
}

func main() {
	v := []int{3, 2, 1}
	var t *Node
	for _, k := range v {
		t = insertNode(t, k)
	}
	printTree(t, "", true)
}
