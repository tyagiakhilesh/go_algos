package algos_test

import (
	algos "algorithms"
	"testing"
)

// Test Initialization of Graph
func TestGraphInitialization(t *testing.T) {
	gr := algos.Graph{}
	gr.Init(true)
	if !gr.Directed {
		t.Errorf("Expected Graph to be Directed")
	}
}

// Test Adding Edges
func TestAddEdge(t *testing.T) {
	var graph algos.Graph = algos.Graph{}
	graph.Init(false)

	// Add edge from 1 to 2 with weight 10
	graph.AddEdge(1, 2, 10)

	if graph.Size != 2 {
		t.Errorf("Expected Number of Edges: 2, Got: %d", graph.Size)
	}
	if graph.Order != 2 {
		t.Errorf("Expected Number of Nodes: 2, Got: %d", graph.Order)
	}

	// Verify the added edge
	edge := graph.AdjacencyList[1].Front().Value.(*algos.EdgeNode)
	if edge.Y != 2 || edge.Weight != 10 {
		t.Errorf("Expected edge to (Y: 2, Weight: 10), Got (Y: %d, Weight: %d)", edge.Y, edge.Weight)
	}
}

// Test Removing Edges
func TestRemoveEdge(t *testing.T) {
	var graph algos.Graph
	graph.Init(false)

	// Add edge (1 -> 2)
	graph.AddEdge(1, 2, 10)

	// Remove the edge
	removed, err, _, _ := graph.RemoveEdge(1, 2)
	if err != nil || !removed {
		t.Errorf("Failed to remove existing edge")
	}

	// Ensure node adjacency list is empty after removal
	if graph.AdjacencyList[1].Len() != 0 {
		t.Errorf("Expected adjacency list for vertex 1 to be empty")
	}

	// Try removing a non-existing edge
	removed, err, _, _ = graph.RemoveEdge(1, 3)
	if removed || err != nil {
		t.Errorf("Expected false and no error when removing a non-existent edge")
	}

	// Remove edge from non-existent node
	removed, err, _, _ = graph.RemoveEdge(3, 4)
	if removed || err == nil {
		t.Errorf("Expected error when removing edge from non-existent node")
	}
}

// Test Directed Graph Functionality
func TestDirectedGraph(t *testing.T) {
	var graph algos.Graph
	graph.Init(true)

	graph.AddEdge(1, 2, 10)
	graph.AddEdge(2, 1, 5)

	if graph.Order != 2 {
		t.Errorf("Expected 2 nodes in graph")
	}

	if graph.Size != 2 {
		t.Errorf("Expected 2 edges in graph")
	}

	// Verify edges in directed graph
	if graph.AdjacencyList[1].Front().Value.(*algos.EdgeNode).Y != 2 {
		t.Errorf("Expected edge from 1 -> 2")
	}

	if graph.AdjacencyList[2].Front().Value.(*algos.EdgeNode).Y != 1 {
		t.Errorf("Expected edge from 2 -> 1")
	}
}
