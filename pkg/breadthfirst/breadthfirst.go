/*
* Route Between Nodes: Given a directed graph, design an algorithm to find out
* whether there is a route between two nodes
 */
package breadthfirst

import (
	. "github.com/Spazzy757/cracking-the-coding-in-go/pkg/graph"
)

func Search(start *Node, end *Node) bool {
	// edge case where start node and end node
	// is the same node
	if start == end {
		return true
	}

	start.State = Visited

	// used as a queue
	var q Queue

	// add starting node to the queue
	q.Push(start)

	var u *Node

	for {
		u = q.Pop()
		for _, v := range u.Adjacent {
			if v.State == Unvisited {
				if v == end {
					return true
				} else {
					v.State = Visiting
					q.Push(v)
				}
			}
		}

		if q.IsEmpty() {
			break
		}
	}
	return false
}
