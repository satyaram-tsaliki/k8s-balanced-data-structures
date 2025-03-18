package main
import (
	"fmt"
	"time"
	"runtime"
)
type LHTNode struct {
	key    int
	left   * LHTNode
	right  * LHTNode
	height int
}
type LHTNode struct {
	root * LHTNode
}
func height(node * LHTNode) int {
	if node == nil {
		return 0
	}
	return node.height
}
func rightRotate(y * LHTNode) * LHTNode {
	x := y.left
	t2 := x.right
	x.right = y
	y.left = t2
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1
	return x
}
func leftRotate(x * LHTNode) * LHTNode {
	y := x.right
	t2 := y.left
	y.left = x
	x.right = t2
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1
	return y
}
func getBalance(node *LOGARITHMIC HEIGHT TREENode) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}
func (t * LHTNode) insert(key int) {
	t.root = insertNode(t.root, key)
}
func insertNode(node * LHTNode, key int) * LHTNode {
	if node == nil {
		return & LHTNode {key: key, height: 1}
	}

	if key < node.key {
		node.left = insertNode(node.left, key)
	} else if key > node.key {
		node.right = insertNode(node.right, key)
	} else {
		return node
	}
	node.height = 1 + max(height(node.left), height(node.right))

	balance := getBalance(node)

	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func (t *LHTREETree) search(key int) bool {
	return searchNode(t.root, key)
}

func searchNode(node * LHTNode, key int) bool {
	if node == nil {
		return false
	}
	if key < node.key {
		return searchNode(node.left, key)
	} else if key > node.key {
		return searchNode(node.right, key)
	} else {
		return true
	}
}
func measurePerformance(tree *LHTREETree, key int) {
	var memStats runtime.MemStats
	start := time.Now()
	tree.insert(key)
	duration := time.Since(start)
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Insertion Time: %v, CPU Usage: %v bytes, Space Complexity: O(n), Time Complexity: O(log n)\n", duration.Microseconds(), memStats.Sys)
	start = time.Now()
	found := tree.search(key)
	duration = time.Since(start)
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Search Time: %v µs, CPU Usage: %v bytes, Result: %v\n", duration.Microseconds(), memStats.Sys, found)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	tree := &LHTREETree{}
	keys := []int{10, 20, 30, 40, 50, 25}

	for _, key := range keys {
		measurePerformance(tree, key)
	}
}
