package main

import "fmt"

type node struct {
	value string
	leaves []node
	end bool
}

func addString (root node, s string) {
	// base case
	if len(s) < 1 {
		return
	}

	var first = s[0]

	// search leaves recursively
	for i := 0; i < len(node.leaves); i++ {
		if node.leaves[i] == first {
			addString(node.leaves[i], s[1:])
			return
		}
	}

	// add needed leaves
	append(root.leaves, node{})


}

func main() {

}