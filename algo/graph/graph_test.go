package graph

import (
	"fmt"
	"testing"
)

func TestNewLinkGraph(t *testing.T) {
	var oj = [][]string{{"a", "b", "d"}, {"b", "a", "c", "d", "f"}, {"c", "b", "f"}, {"d", "a", "b", "f"}, {"f", "b", "c", "d"}}
	x := NewLinkGraph(oj)
	for i, i2 := range x {
		fmt.Println(i)
		fmt.Printf("%+v", i2)
		fmt.Println()
	}
}

func TestNewGraph(t *testing.T) {
	NewGraph()
}

func TestGraph_AddEdge(t *testing.T) {
	g := NewGraph()
	g.AddEdge("0", "1", 5)
	g.AddEdge("0", "5", 2)
	g.AddEdge("1", "2", 4)
	g.AddEdge("2", "3", 9)
	g.AddEdge("3", "4", 7)
	g.AddEdge("3", "5", 3)
	g.AddEdge("3", "5", 3)
	g.AddEdge("5", "4", 8)
	g.AddEdge("5", "2", 1)
	for _, v := range g.GetVertices() {
		fmt.Println(v.String())
		//for _, x := range v.GetConnections() {
		//	fmt.Println("<", x.start, x.end, ">", x.weight)
		//}
	}
}

func TestBFS(t *testing.T) {
	g := NewGraph()
	g.AddEdge("0", "1", 5)
	g.AddEdge("0", "5", 2)
	g.AddEdge("1", "2", 4)
	g.AddEdge("2", "3", 9)
	g.AddEdge("3", "4", 7)
	g.AddEdge("3", "5", 3)
	g.AddEdge("3", "5", 3)
	g.AddEdge("5", "4", 8)
	g.AddEdge("5", "2", 1)
	BFS(g, "0")
	for k, v := range g.VertexList {
		fmt.Printf("%s->color:%d,distanct:%d [%d,%d] \n", k, v.Color, v.Distance, v.First, v.Last)
	}

}

func TestDFS(t *testing.T) {
	g := NewGraph()
	g.AddEdge("0", "1", 5)
	g.AddEdge("0", "5", 2)
	g.AddEdge("1", "2", 4)
	g.AddEdge("2", "3", 9)
	g.AddEdge("3", "4", 7)
	g.AddEdge("3", "5", 3)
	g.AddEdge("3", "5", 3)
	g.AddEdge("5", "4", 8)
	g.AddEdge("5", "2", 1)
	DFS(g, "4")
	for k, v := range g.VertexList {
		fmt.Printf("%s->color:%d,distanct:%d [%d,%d] \n", k, v.Color, v.Distance, v.First, v.Last)
	}
}
