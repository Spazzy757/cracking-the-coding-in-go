package graph

type State int

const (
	Unvisited State = iota
	Visited
	Visiting
)

type Node struct {
	Id       string
	State    State
	Adjacent []*Node
}
