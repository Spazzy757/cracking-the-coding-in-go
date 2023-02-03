package breadthfirst

import (
	"testing"

	. "github.com/Spazzy757/cracking-the-coding-in-go/pkg/graph"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	/*
	* Test case will use the following Graph
	*┌──────┐
	*│A     │
	*└┬─┬──┬┘
	* │┌▽┐┌▽┐
	* ││B││D│
	*	│└┬┘└┬┘
	*┌▽─▽┐┌▽┐
	*│C  ││F│
	*└───┘└┬┘
	*┌─────▽┐
	*│G     │
	*└──────┘
	 */
	g := &Node{Id: "G"}
	f := &Node{Id: "F", Adjacent: []*Node{g}}
	d := &Node{Id: "D", Adjacent: []*Node{f}}
	c := &Node{Id: "C"}
	b := &Node{Id: "B", Adjacent: []*Node{c}}
	a := &Node{Id: "A", Adjacent: []*Node{b, c, d}}
	assert := require.New(t)
	t.Run("nodes are connected", func(t *testing.T) {
		// A is connected to G
		c := Search(a, g)
		assert.True(c)
	})
	t.Run("nodes are not connected", func(t *testing.T) {
		// A is connected to G
		c := Search(g, b)
		assert.False(c)
	})
	t.Run("nodes are the same", func(t *testing.T) {
		// A is connected to G
		c := Search(b, b)
		assert.True(c)
	})
}
