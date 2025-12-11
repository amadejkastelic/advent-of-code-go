package graph

type Node struct {
	ID       string
	Adjacent []*Node
}

func (n *Node) TraverseFunc(fn func(*Node)) {
	fn(n)
	for _, neighbor := range n.Adjacent {
		neighbor.TraverseFunc(fn)
	}
}

func (n *Node) TraverseFuncWithPath(fn func(*Node, map[string]bool), visited map[string]bool) {
	visited[n.ID] = true
	fn(n, visited)
	for _, neighbor := range n.Adjacent {
		if !visited[neighbor.ID] {
			neighbor.TraverseFuncWithPath(fn, visited)
		}
	}
	delete(visited, n.ID)
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(id string) {
	if _, exists := g.Nodes[id]; !exists {
		g.Nodes[id] = &Node{
			ID:       id,
			Adjacent: []*Node{},
		}
	}
}

func (g *Graph) AddEdge(fromID, toID string) {
	g.AddNode(fromID)
	g.AddNode(toID)

	fromNode := g.Nodes[fromID]
	toNode := g.Nodes[toID]

	fromNode.Adjacent = append(fromNode.Adjacent, toNode)
}

func (g *Graph) GetNode(id string) (*Node, bool) {
	node, exists := g.Nodes[id]
	return node, exists
}
