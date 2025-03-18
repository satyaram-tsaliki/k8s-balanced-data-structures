import (
	"fmt"
	"runtime"
	"time"
)

const t = 2

type BTreeNode struct {
	keys     []int
	children []*BTreeNode
	leaf     bool
	n        int
}

type BTree struct {
	root *BTreeNode
}

func newBTreeNode(leaf bool) *BTreeNode {
	return &BTreeNode{leaf: leaf, keys: make([]int, 2*t-1), children: make([]*BTreeNode, 2*t), n: 0}
}

func (tree *BTree) insert(key int) {
	if tree.root == nil {
		tree.root = newBTreeNode(true)
		tree.root.keys[0] = key
		tree.root.n = 1
	} else {
		if tree.root.n == 2*t-1 {
			newRoot := newBTreeNode(false)
			newRoot.children[0] = tree.root
			splitChild(newRoot, 0, tree.root)
			tree.root = newRoot
		}
		insertNonFull(tree.root, key)
	}
}

func splitChild(parent *BTreeNode, i int, fullChild *BTreeNode) {
	newNode := newBTreeNode(fullChild.leaf)
	newNode.n = t - 1
	for j := 0; j < t-1; j++ {
		newNode.keys[j] = fullChild.keys[j+t]
	}
	if !fullChild.leaf {
		for j := 0; j < t; j++ {
			newNode.children[j] = fullChild.children[j+t]
		}
	}
	fullChild.n = t - 1
	for j := parent.n; j >= i+1; j-- {
		parent.children[j+1] = parent.children[j]
	}
	parent.children[i+1] = newNode
	for j := parent.n - 1; j >= i; j-- {
		parent.keys[j+1] = parent.keys[j]
	}
	parent.keys[i] = fullChild.keys[t-1]
	parent.n++
}

func insertNonFull(node *BTreeNode, key int) {
	i := node.n - 1
	if node.leaf {
		for i >= 0 && key < node.keys[i] {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = key
		node.n++
	} else {
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++
		if node.children[i].n == 2*t-1 {
			splitChild(node, i, node.children[i])
			if key > node.keys[i] {
				i++
			}
		}
		insertNonFull(node.children[i], key)
	}
}

func measureBTreePerformance(tree *BTree, key int) {
	var memStats runtime.MemStats

	
	start := time.Now()
	tree.insert(key)
	duration := time.Since(start)
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Insertion Time: %v µs, CPU Usage: %v bytes, Space Complexity: O(n), Time Complexity: O(log n)\n", duration.Microseconds(), memStats.Sys)
}

func main() {
	tree := &BTree{}
	keys := []int{10, 20, 30, 40, 50, 25}

	for _, key := range keys {
		measureBTreePerformance(tree, key)
	}
}
