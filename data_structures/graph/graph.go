package graph

import (
	"errors"
)

type VertexId uint

type Vertices []VertexId

type Edge struct {
	From VertexId
	To   VertexId
}

type Graph struct {
	edges      map[VertexId]map[VertexId]int
	edgesCount int
	isDirected bool
}

type EdgesIterable interface {
	EdgesIter() <-chan Edge
}

type VerticesIterable interface {
	VerticesIter() <-chan VertexId
}

func NewDirected() *Graph {
	return &Graph{
		edges:      make(map[VertexId]map[VertexId]int),
		edgesCount: 0,
		isDirected: true,
	}
}
func NewUndirected() *Graph {
	return &Graph{
		edges:      make(map[VertexId]map[VertexId]int),
		edgesCount: 0,
		isDirected: false,
	}
}

func (g *Graph) EdgesCount() int {
	return g.edgesCount
}

func (g *Graph) VerticesCount() int {
	return len(g.edges)
}

func (g *Graph) IsDirected() bool {
	return g.isDirected
}

func (g *Graph) EdgesIter() <-chan Edge {
	ch := make(chan Edge)
	go func() {
		for from, connectedVertices := range g.edges {
			for to, _ := range connectedVertices {
				if g.IsDirected() {
					ch <- Edge{from, to}
				} else {
					if from < to {
						ch <- Edge{from, to}
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func (g *Graph) VerticesIter() <-chan VertexId {
	ch := make(chan VertexId)
	go func() {
		for vertex, _ := range g.edges {
			ch <- vertex
		}
		close(ch)
	}()
	return ch
}

func (g *Graph) CheckVertex(vertex VertexId) bool {
	_, exists := g.edges[vertex]
	return exists
}

func (g *Graph) TouchVertex(v VertexId) {
	if _, ok := g.edges[v]; !ok {
		g.edges[v] = make(map[VertexId]int)
	}
}

func (g *Graph) AddVertex(v VertexId) error {
	i, _ := g.edges[v]
	if i != nil {
		return errors.New("Vertex is already exist")
	}
	g.edges[v] = make(map[VertexId]int)
	return nil
}

func (g *Graph) RemoveVertex(v VertexId) error {
	if !g.IsVertex(v) {
		return errors.New("Unknow vertex")
	}

	delete(g.edges, v)
	for _, connectedVertices := range g.edges {
		delete(connectedVertices, v)
	}
	return nil
}

func (g *Graph) IsVertex(v VertexId) (exist bool) {
	_, exist = g.edges[v]
	return
}

func (g *Graph) AddEdge(from, to VertexId, weight int) error {
	if from == to {
		return errors.New("Cannot add self loop")
	}

	if !g.CheckVertex(from) || !g.CheckVertex(to) {
		return errors.New("Vertices don't exist")
	}

	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]
	if i > 0 || j > 0 {
		return errors.New("Edge already added")
	}
	g.TouchVertex(from)
	g.TouchVertex(to)

	g.edges[from][to] = weight
	if !g.IsDirected() {
		g.edges[to][from] = weight
	}
	g.edgesCount += 1
	return nil
}

func (g *Graph) RemoveEdge(from, to VertexId) error {
	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]

	if i == -1 || j == -1 {
		return errors.New("Edge doesn't exist")
	}

	g.edges[from][to] = -1
	if !g.IsDirected() {
		g.edges[to][from] = -1
	}
	g.edgesCount -= 1

	return nil
}

func (g *Graph) IsEdge(from, to VertexId) bool {
	connected, ok := g.edges[from]
	if !ok {
		return false
	}
	weight := connected[to]
	return weight > 0
}

func (g *Graph) Order() int {
	return len(g.edges)
}

func (g *Graph) GetEdge(from, to VertexId) int {
	return g.edges[from][to]
}

func (g *Graph) GetNeighbours(v VertexId) VerticesIterable {
	iterator := func() <-chan VertexId {
		ch := make(chan VertexId)
		go func() {
			if connected, ok := g.edges[v]; ok {
				for VertexId, _ := range connected {
					ch <- VertexId
				}
			}
			close(ch)
		}()
		return ch
	}
	return VerticesIterable(&vertexIterableHelper{iterFunc: iterator})
}

type vertexIterableHelper struct {
	iterFunc func() <-chan VertexId
}

func (helper *vertexIterableHelper) VerticesIter() <-chan VertexId {
	return helper.iterFunc()
}
