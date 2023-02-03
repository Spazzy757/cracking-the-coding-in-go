/*
* Route Between Nodes: Given a directed graph, design an algorithm to find out
* whether there is a route between two nodes
 */
package breadthfirst

type State int

const (
	Unvisited State = iota
	Visited
	Visiting
)

type Node struct {
	id       string
	State    string
	Adjacent []Node
}

type Queue struct {
	List[Node]
}

func (q *Queue) IsEmpty() {
	return len(q.List) == 0
}

func (q *Queue) Append(n Node) {
	q.List = append(q.List, n)
}

func (q *Queue) Pop() Node {
	if q.IsEmpty() {
		panic("cant pop an empty list")
	}

	u := q.List[0]
	q.List = q.List[1:]

	return u
}

func Search(start Node, end Node) Bool {
	if start == end {
		return true
	}

	// used as a queue
	var q Queue

	start.State
	q = q.Append(start)

	var u Node

	for {
		u = q.Pop()
		for _, v := range u.Adjacent {
			if v.State == Unvisited {
				if v == end {
					return true
				} else {
					v.State = Visting
					q.Append(v)
				}
			}
		}

		if q.IsEmpty() {
			break
		}
	}
	return false
}
