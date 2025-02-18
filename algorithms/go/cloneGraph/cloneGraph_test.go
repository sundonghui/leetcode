package clonegraph

import (
	"reflect"
	"testing"
)

func Test_cloneGraph_example1(t *testing.T) {
	node := &Node{Val: 1, Neighbors: []*Node{{Val: 2, Neighbors: []*Node{{Val: 3}}}, {Val: 4}}}
	cloneGraph(node)
	if !reflect.DeepEqual(cloneGraph(node), node) {
		t.Errorf("cloneGraph() = %v, want %v", cloneGraph(node), node)
	}
}
