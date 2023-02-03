package graph

type Queue struct {
	list []*Node
}

func (q *Queue) IsEmpty() bool {
	return len(q.list) == 0
}

func (q *Queue) Push(n *Node) {
	q.list = append(q.list, n)
}

func (q *Queue) Pop() *Node {
	if q.IsEmpty() {
		panic("cant pop an empty list")
	}

	u := q.list[0]
	q.list = q.list[1:]

	return u
}
