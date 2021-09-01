package graph

import (
	"fmt"
	"oi.io/algo/queue"
	"strings"
)

type VertexInf interface {
	AddNeighbor(VertexID string, weight int)
	String() string
	GetConnections() map[string]Edge
	GetID() string
	GetWeight(VertexID string) int
	SetColor(int)
	GetColor() int
	SetDistance(int)
	GetDistance() int
}

type Edge struct {
	start  string
	end    string
	weight int
}

type Vertex struct {
	ID        string
	Connected map[string]Edge
	Distance  int
	Color     int
	First     int
	Last      int
	Pred      *Vertex
}

func NewVertex(VertexID string) *Vertex {
	return &Vertex{
		ID:        VertexID,
		Connected: make(map[string]Edge),
	}
}

func (v *Vertex) GetColor() int {
	return v.Color
}

func (v *Vertex) SetColor(color int) {
	v.Color = color
}

func (v *Vertex) GetDistance() int {
	return v.Distance
}

func (v *Vertex) SetDistance(distance int) {
	v.Distance = distance
}

func (v *Vertex) SetPred(p *Vertex) {
	v.Pred = p
}

func (v *Vertex) GetPred(p *Vertex) {
	v.Pred = p
}

func (v *Vertex) AddNeighbor(VertexID string, weight int) {
	if v.Connected == nil {
		v.Connected = make(map[string]Edge)
	}
	v.Connected[VertexID] = Edge{
		start:  v.GetID(),
		end:    VertexID,
		weight: weight,
	}
	return
}

func (v *Vertex) String() string {
	var keys []string
	for k, _ := range v.GetConnections() {
		keys = append(keys, k)
	}
	key := strings.Join(keys, ",")
	return fmt.Sprintf("Vertex %s connect with (%s)", v.ID, key)
}

func (v *Vertex) GetConnections() map[string]Edge {
	return v.Connected
}

func (v *Vertex) GetID() string {
	return v.ID
}

func (v *Vertex) GetWeight(VertexID string) int {
	return v.Connected[VertexID].weight //todo fixed not exist
}

type GraphInf interface {
	AddVertex(vertexID string)
	GetVertex(vertexID string) *Vertex
	AddEdge(start string, end string, weight int)
	GetVertices() []string
}

type Graph struct {
	VertexList  map[string]*Vertex
	NumOfVertex int
}

func NewGraph() *Graph {
	return &Graph{
		VertexList:  make(map[string]*Vertex),
		NumOfVertex: 0,
	}
}

func (g *Graph) AddVertex(vertexID string) {
	if g.VertexList == nil {
		g.VertexList = make(map[string]*Vertex)
	}
	g.VertexList[vertexID] = NewVertex(vertexID)
	g.NumOfVertex += 1
}

func (g *Graph) GetVertex(vertexID string) *Vertex {
	return g.VertexList[vertexID]
}

func (g *Graph) AddEdge(start string, end string, weight int) {
	var vertex1, vertex2 *Vertex
	vertex1, ok := g.VertexList[start]
	if !ok {
		vertex1 = NewVertex(start)
	}
	vertex2, ok = g.VertexList[end]
	if !ok {
		vertex2 = NewVertex(end)
		g.VertexList[end] = vertex2
	}
	vertex1.AddNeighbor(end, weight)
	g.VertexList[start] = vertex1
}

func (g *Graph) GetVertices() (ret []*Vertex) {
	for _, v := range g.VertexList {
		ret = append(ret, v)
	}
	return
}

type LinkGraph struct {
	name string
	next *LinkGraph
}

func (g *LinkGraph) Add(v string) {
	n := &LinkGraph{name: v}
	last := g
	for {
		if last.next != nil {
			last = last.next
		} else {
			last.next = n
			break
		}
	}
}

func New(v []string) *LinkGraph {
	var ret *LinkGraph
	for _, name := range v {
		n := &LinkGraph{name: name}
		if ret == nil {
			ret = n
		} else {
			ret.Add(name)
		}
	}
	return ret
}

func NewLinkGraph(r [][]string) map[string]*LinkGraph {
	mg := make(map[string]*LinkGraph)
	for _, arr := range r {
		mg[arr[0]] = New(arr)
	}
	return mg
}

const (
	White = iota
	Gray
	Black
)

func BFS(g *Graph, start string) {
	s := g.GetVertex(start)
	s.SetDistance(0)
	var q = queue.NewQueue()
	count := 0
	q.EnQueue(s)
	for !q.IsEmpty() {
		cur := q.DeQueue()
		x := cur.(*Vertex)
		x.First = count
		for k, _ := range x.GetConnections() {
			row := g.GetVertex(k)
			if row == nil {
				fmt.Printf("get %s with nil\n", k)
				continue
			}
			count += 1
			if row.GetColor() == White {
				row.SetDistance(x.Distance + 1)
				row.SetColor(Gray)
				q.EnQueue(row)
			}
		}
		x.SetColor(Black)
		x.Last = count
	}
}

func DFS(g *Graph, start string) {
	s := g.GetVertex(start)
	s.SetDistance(0)
	var q = queue.NewQueue()
	count := 0
	q.EnQueue(s)
	for !q.IsEmpty() {
		cur := q.DeQueue()
		x := cur.(*Vertex)
		x.First = count
		for k, _ := range x.GetConnections() {
			row := g.GetVertex(k)
			if row == nil {
				fmt.Printf("get %s with nil\n", k)
				continue
			}
			count += 1
			if row.GetColor() == White {
				row.SetDistance(x.Distance + 1)
				row.SetColor(Gray)
				q.EnQueue(row)
			}
		}
		x.SetColor(Black)
		x.Last = count
	}
}
