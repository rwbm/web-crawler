package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
}

func TestAddNodeThenHasNode(t *testing.T) {
	label := "test"
	g := NewGraph()
	g.AddNode(label)
	hasNode := g.HasNode(label)

	assert.Equal(t, true, hasNode)
}

func TestHasNodeWithNoNode(t *testing.T) {
	label := "test"
	g := NewGraph()
	hasNode := g.HasNode(label)

	assert.Equal(t, false, hasNode)
}

func TestHasNodeWithWithDuplicateAddNodes(t *testing.T) {
	label := "test"
	g := NewGraph()
	g.AddNode(label)
	g.AddNode(label)
	hasNode := g.HasNode(label)

	assert.Equal(t, true, hasNode)
}

func TestAddEdgeWithValidNodes(t *testing.T) {
	label1 := "node1"
	label2 := "node2"
	g := NewGraph()
	g.AddNode(label1)
	g.AddNode(label2)
	err := g.AddEdge(label1, label2)

	assert.NoError(t, err)
}

func TestAddEdgeWithoutAddingFromNode(t *testing.T) {
	label1 := "node1"
	label2 := "node2"
	g := NewGraph()
	g.AddNode(label2)
	err := g.AddEdge(label1, label2)

	assert.Error(t, err)
}

func TestAddEdgeWithoutAddingToNode(t *testing.T) {
	label1 := "node1"
	label2 := "node2"
	g := NewGraph()
	g.AddNode(label1)
	err := g.AddEdge(label1, label2)

	assert.Error(t, err)
}

func TestAddEdgeWithoutAddingAnyNode(t *testing.T) {
	label1 := "node1"
	label2 := "node2"
	g := NewGraph()
	err := g.AddEdge(label1, label2)

	assert.Error(t, err)
}

func TestBFSWithNoEdges(t *testing.T) {
	var edges [][]string

	f := func(from string, to string) {
		testPair := []string{from, to}
		edges = append(edges, testPair)
	}

	label1 := "node1"
	label2 := "node2"
	g := NewGraph()
	g.AddNode(label1)
	g.AddNode(label2)

	err := g.BFS(label1, f)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(edges))
}

func TestBFSWithNoNodes(t *testing.T) {
	var edges [][]string

	f := func(from string, to string) {
		testPair := []string{from, to}
		edges = append(edges, testPair)
	}

	label1 := "node1"

	g := NewGraph()
	err := g.BFS(label1, f)

	assert.Error(t, err)
}

func TestBFSWithUnknownNode(t *testing.T) {
	var edges [][]string

	f := func(from string, to string) {
		testPair := []string{from, to}
		edges = append(edges, testPair)
	}

	label1 := "node1"
	label2 := "node2"
	g := NewGraph()
	g.AddNode(label1)

	err := g.BFS(label2, f)

	assert.Error(t, err)
}
