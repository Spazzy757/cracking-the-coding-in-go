/*
* Route Between Nodes: Given a directed graph, design an algorithm to find out
* whether there is a route between two nodes
 */
package breadthfirst

type Node struct {
	id       string
	state    string
	adjacent []Node
}

func search(start Node, end Node) Bool {
	if start == end {
		return true
	}

	var nodeList []Node
}
