package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	key   string
}

func createNode(key string) *Node {
	return &Node{nil, nil, key}
}

func preOrder(root *Node) string {
	if root == nil {
		return ""
	}

	return root.key + preOrder(root.left) + preOrder(root.right)
}

func postOrder(root *Node) string {
	if root == nil {
		return ""
	}
	return postOrder(root.left) + postOrder(root.right) + root.key
}

func main() {
	root := createNode("+")
	root.left = createNode("a")
	root.right = createNode("-")
	root.right.left = createNode("b")
	root.right.right = createNode("c")

	fmt.Println("Preorder: ", preOrder(root))
	fmt.Println("Postorder: ", postOrder(root))

}
