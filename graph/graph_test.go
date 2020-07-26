package graph

import (
	"testing"
)

func TestBuildGraph(t *testing.T) {
	g := NewGraph(5, false)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(5, 4)
	g.AddEdge(2, 4)
	g.AddEdge(2, 3)
	g.AddEdge(4, 3)

	if g.m != 7 {
		t.Error("Invalid number of edges")
	}

	g.Display()
}

func TestCheckEdge(t *testing.T) {
	g := NewGraph(5, false)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(5, 4)
	g.AddEdge(2, 4)
	g.AddEdge(2, 3)
	g.AddEdge(4, 3)

	if g.CheckEdge(5, 4) == false {
		t.Error("Incorrect check edge")
	}

	if g.CheckEdge(1, 4) == true {
		t.Error("Incorrect check edge")
	}
}
