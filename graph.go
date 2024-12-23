package algos

import (
	"container/list"
	"fmt"
)

type EdgeNode struct {
	Y      int
	Weight int
	Next   *EdgeNode
}

type Graph struct {
	// Number of edges
	Size int
	// Number of vertices
	Order int
	// If Graph is directed or not
	Directed bool
	// Degree of the graph
	Degree        int
	AdjacencyList map[int]*list.List
}

func (g *Graph) Init(IsDirected bool) {
	g.Directed = IsDirected
	g.AdjacencyList = make(map[int]*list.List)
}

func (g *Graph) AddEdge(x, y int, weight int) {
	addEdge(x, y, weight, g)
	if !g.Directed {
		addEdge(y, x, weight, g)
	}
}

func addEdge(x int, y int, weight int, g *Graph) {
	if nil == g.AdjacencyList[x] {
		g.AdjacencyList[x] = list.New()
		g.Order++
	}
	g.AdjacencyList[x].PushBack(&EdgeNode{Y: y, Weight: weight})
	g.Size++
}

func (g *Graph) RemoveEdge(x, y int) (bool, error, bool, error) {
	r1, e1 := removeEdge(x, y, g)
	var r2 = false
	var e2 error = nil
	if !g.Directed {
		r2, e2 = removeEdge(y, x, g)
	}
	return r1, e1, r2, e2
}

func removeEdge(x int, y int, g *Graph) (bool, error) {
	if nil != g.AdjacencyList[x] {
		for e := g.AdjacencyList[x].Front(); e != nil; e = e.Next() {
			if e.Value.(*EdgeNode).Y == y {
				g.AdjacencyList[x].Remove(e)
				g.Size--
				return true, nil
			}
		}
		return false, nil
	} else {
		return false, fmt.Errorf(`node x: %d doesn't exist'`, x)
	}
}
