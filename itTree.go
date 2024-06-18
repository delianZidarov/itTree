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
			node.height = newHeight(node)
			balance := height(node.left) - height(node.right)
			switch {
			case balance > 1:
				// base case: if node.left.right is nil then the nodes are
				//lined up and only one rotation is necessary. Using the height
				//of the branches generalizes the base case
				if height(node.left.left) > height(node.left.right) {
					if i == 0 {
						//if your are rotating the root node set the return to 
						//the new root after rotating
						n = rotateRight(node)
					} else if i > 0 {
						//if your are not at the root node go one level up and 
						//set the previous left || right node to the new one
						if v[i-1].left == node {
							v[i-1].left = rotateRight(node)
						} else if v[i-1].right == node {
							v[i-1].right = rotateRight(node)
						}
					}
				// base case: if node.left.left is nil then the nodes are
				//stagared and two rotations are necessary. Using the height
				//of the branches generalizes the base case
				} else if height(node.left.left) < height(node.left.right) {
					node.left = rotateLeft(node.left)
					if i == 0 {
						n = rotateRight(node)
					} else if i > 0 {
						if v[i-1].left == node {
							v[i-1].left = rotateRight(node)
						} else if v[i-1].right == node {
							v[i-1].right = rotateRight(node)
						}
					}
				}

			case balance < -1:
				if height(node.right.right) > height(node.right.left) {
					if i == 0 {
						n = rotateLeft(node)
					} else if i > 0 {
						if v[i-1].left == node {
							v[i-1].left = rotateLeft(node)
						} else if v[i-1].right == node {
							v[i-1].right = rotateLeft(node)
						}
					}
				} else if height(node.right.right) < height(node.right.left) {
					node.right = rotateRight(node.right)
					if i == 0 {
						n = rotateLeft(node)
					} else if i > 0 {
						if v[i-1].left == node {
							v[i-1].left = rotateLeft(node)
						} else if v[i-1].right == node {
							v[i-1].right = rotateLeft(node)
						}
					}
				}
			}
		}
	}
	return n
}

func rotateRight(n *Node) *Node {
	t := n.left
	t1 := t.right
	t.right = n
	n.left = t1
	n.height = newHeight(n)
	t.height = newHeight(t)
	return t
}

func rotateLeft(n *Node) *Node {
	t := n.right
	t1 := n.right.left
	t.left = n
	n.right = t1
	n.height = newHeight(n)
	t.height = newHeight(t)
	return t
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

func printNode(n *Node) {
	v := make([]*Node, 0)
	v = append(v, n)

	var bWT func(n *Node)
	bWT = func(n *Node) {
		if n == nil {
			return
		}
		if n.left != nil {
			v = append(v, n.left)
		}
		if n.right != nil {
			v = append(v, n.right)
		}
		bWT(n.left)
		bWT(n.right)
	}
	bWT(n)

	for _, k := range v {
		fmt.Println(k.key)
	}
}

func main() {
	v := []int{1, 5, 8, 6, 9, 7,10,11,12,13,14}
	var t *Node
	for _, k := range v {
		t = insertNode(t, k)
	}
	printTree(t, "", true)
	// printNode(t)
}
