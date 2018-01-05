package httprouter

import "fmt"

// WalkPrint 打印整个 tree 的结构
func WalkPrint(r *Router, method string) {
	tree := r.trees[method]
	if tree != nil {
		walkAndPrintNodeRecursive(tree, 1)
	}
}

func walkAndPrintNodeRecursive(n *node, level int) {
	if n == nil {
		return
	}

	fmt.Printf("node: %#v ; level: %d\n", n, level)
	for i := 0; i < len(n.children); i++ {
		walkAndPrintNodeRecursive(n.children[i], level+1)
	}
}
