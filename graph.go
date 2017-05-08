package fabric

// Graph can be either UI DDAG, Temporal DAG or VDG
type Graph struct {
	Nodes []int
	Edges map[int][]int // each node has a list of nodes that it points too
}

// TODO: should this function be passed a list of node int IDs generated by the user
// OR should the user be expected to just do a myGraph.Nodes = myList after creating an empty graph (??)
// What will be some techniques for generating a graph via partial orderings (?)
func NewGraph() *Graph {
	var nodes []int
	return &Graph{
		Nodes: nodes,
		Edges: make(map[int][]int),
	}
}

func (g *Graph) CycleDetect(start int) bool {
	var seen []int
	var done []int

	return g.cycleDfs(start, seen, done)
}

func (g *Graph) GetAdjacents(node int) []int {
	return g.Edges[node]
}

/* Recursive Depth-First-Search for Cycle Detection*/
func (g *Graph) cycleDfs(start int, seen, done []int) bool {

	seen = append(seen, start)
	adj := g.Edges[start]
	for _, v := range adj {
		if contains(done, v) {
			continue
		}

		if contains(seen, i) {
			return true
		}

		if g.cycleDfs(v, seen, done) {
			return true
		}
	}
	seen = seen[:len(seen)-1]
	done = append(done, start)
	return false
}

// NOTE: this method cannot be used on UI DDAGs
func (g *Graph) RemoveNode(node int) {
	// TODO:
	// remove node from node list
	// remove all of nodes edges from edges map
}

func (g *Graph) AppendNode(node int) {
	// TODO: add node to nodes list
}

// AppendEdge adds an edge that points from dependent to dependency
func (g *Graph) AppendEdge(source, dest int) {
	// TODO: add an edge to edges map which is source dependent on destination
}

func (g *Graph) RemoveEdge(source, dest int) {
	// TODO: removes the edge between source and dest in the edges map
}

func (g *Graph) AppendSubGraph(graph Graph) {
	// TODO:
	// add all nodes in new graph to existing graph
	// add all edges in new graph to existing graph
	// cycle detection on new total graph
}

// NOTE: this method cannot be used on UI DDAGs
func (g *Graph) RemoveSubGraph(nodes []int) {
	// TODO: removes list of nodes and all their edges (that they are sources for)
}
