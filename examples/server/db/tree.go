package db

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JKhawaja/fabric"
)

// Tree ...
type Tree struct {
	Root     *TreeNode
	Sections fabric.NodeList
	Nodes    fabric.NodeList
	Edges    fabric.EdgeList
}

// NewSection takes a session id and creates a root node for a branch section
// that will be dedicated to that session (here session ids are behaving like user ids)
func (t *Tree) NewSection(id int) *fabric.Node {
	// create section node (the value will be the session id)
	n := TreeNode{
		Id:    t.GenNodeID(),
		Value: id,
	}

	var i interface{} = n
	in := i.(fabric.Node)

	t.Sections = append(t.Sections, &in)

	// create edge from root to section node
	e := TreeEdge{
		Id:          t.GenEdgeID(),
		Source:      t.Root,
		Destination: &n,
	}
	var inf interface{} = e
	ie := inf.(fabric.Edge)
	t.Edges = append(t.Edges, &ie)

	return &in
}

func containsNode(l fabric.NodeList, id int) bool {
	for _, vp := range l {
		v := *vp
		if v.ID() == id {
			return true
		}
	}
	return false
}

func containsEdge(l fabric.EdgeList, id int) bool {
	for _, vp := range l {
		v := *vp
		if v.ID() == id {
			return true
		}
	}
	return false
}

// CreateNode will create a node and add it to the section and tree data store
func (t *Tree) CreateNode(sp *fabric.Section, value interface{}) (*TreeNode, error) {

	np := CreateNode(t, value)
	n := *np

	// update section with new node
	s := *sp
	nodes := *s.ListNodes()
	var i interface{} = n
	in := i.(fabric.Node)
	nodes = append(nodes, &in)
	s.UpdateNodeList(&nodes)

	return np, nil
}

// TODO: this function will iterate through the sections node list twice,
// rewrite function so it does everything it needs to on a single iteration of
// the sections nodelist.
func (t *Tree) RemoveNode(sp *fabric.Section, id int) error {
	// verify that node is in section before being removed
	s := *sp
	nlp := s.ListNodes()
	nodes := *nlp
	if containsNode(nodes, id) {
		RemoveNode(t, id)
	} else {
		return fmt.Errorf("Node is not in section. Cannot remove.")
	}

	// remove node from section list and update section with new list
	for i, np := range nodes {
		n := *np
		if n.ID() == id {
			nodes = append(nodes[:i], nodes[i+1:]...)
			break
		}
	}

	s.UpdateNodeList(&nodes)

	return nil
}

// CreateEdge ...
func (t *Tree) CreateEdge(sp *fabric.Section, n1, n2 *TreeNode) (*TreeEdge, error) {
	ep := CreateEdge(t, n1, n2)

	// update section with new edge
	s := *sp
	elp := s.ListEdges()
	edges := *elp
	e := *ep
	var i interface{} = e
	ie := i.(fabric.Edge)
	edges = append(edges, &ie)
	s.UpdateEdgeList(&edges)

	return ep, nil
}

// TODO: this function will iterate through the sections edge list twice,
// rewrite function so it does everything it needs to on a single iteration of
// the sections edgelist.
func (t *Tree) RemoveEdge(sp *fabric.Section, id int) error {
	// verify that edge is in section before being removed
	s := *sp
	elp := s.ListEdges()
	edges := *elp
	if containsEdge(edges, id) {
		RemoveEdge(t, id)
	} else {
		return fmt.Errorf("Edge is not in section. Cannot remove.")
	}

	// remove edge from section list and update section with new list
	for i, ep := range edges {
		e := *ep
		if e.ID() == id {
			edges = append(edges[:i], edges[i+1:]...)
			break
		}
	}

	s.UpdateEdgeList(&edges)

	return nil
}

// ReadNodeValue ...
func (t *Tree) ReadNodeValue(sp *fabric.Section, id int) (interface{}, error) {
	// verify that node is in section before being read
	s := *sp
	var value interface{}
	nlp := s.ListNodes()
	nodes := *nlp
	if containsNode(nodes, id) {
		value = ReadNodeValue(t, id)
	} else {
		return value, fmt.Errorf("Node is not in section. Cannot read value.")
	}

	return value, nil
}

func (t *Tree) UpdateNodeValue(sp *fabric.Section, id int, value interface{}) error {
	// verify that node is in section before being updated
	s := *sp
	nlp := s.ListNodes()
	nodes := *nlp
	if containsNode(nodes, id) {
		UpdateNodeValue(t, id, value)
	} else {
		return fmt.Errorf("Node is not in section. Cannot update value.")
	}

	return nil
}

// NewTree ...
func NewTree() *Tree {
	t := Tree{}
	n := &TreeNode{
		Id: t.GenNodeID(),
	}
	t.Root = n

	var i interface{} = n
	in := i.(fabric.Node)
	var nl fabric.NodeList
	nl = append(nl, &in)
	t.Nodes = nl

	el := make(fabric.EdgeList, 0)
	t.Edges = el

	return &t
}

// GenNodeID ...
// Generate an ID for a CDS Node
func (t Tree) GenNodeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, np := range t.Nodes {
		n := *np
		if n.ID() == id {
			id = t.GenNodeID()
		}
	}

	return id
}

// GenEdgeID ...
// Generate an ID for a CDS Edge
func (t Tree) GenEdgeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, ep := range t.Edges {
		e := *ep
		if e.ID() == id {
			id = t.GenEdgeID()
		}
	}

	return id

}

// ListNodes ...
func (t Tree) ListNodes() fabric.NodeList {
	return t.Nodes
}

// ListEdges ...
func (t Tree) ListEdges() fabric.EdgeList {
	return t.Edges
}

// TreeNode satisfies fabric.Node interface
type TreeNode struct {
	Id    int
	Value interface{}
	Imm   bool
}

// ID ...
func (t TreeNode) ID() int {
	return t.Id
}

// Immutable ...
func (t TreeNode) Immutable() bool {
	return t.Imm
}

// TreeEdge satisfies the fabric.Edge interface
type TreeEdge struct {
	Id          int
	Source      *TreeNode
	Destination *TreeNode
	Imm         bool
}

// ID ...
func (t TreeEdge) ID() int {
	return t.Id
}

// GetSource ...
func (t TreeEdge) GetSource() *fabric.Node {
	var i interface{} = *t.Source
	in := i.(fabric.Node)
	return &in
}

// GetDestination ...
func (t TreeEdge) GetDestination() *fabric.Node {
	var i interface{} = *t.Destination
	in := i.(fabric.Node)
	return &in
}

// Immutable ...
func (t TreeEdge) Immutable() bool {
	return t.Imm
}