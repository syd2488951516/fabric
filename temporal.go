package fabric

// This is what is assigned to Threads beyond the first thread assigned to a particular UI
type Temporal interface {
	DGNode
	Root() UI
	Dependencies() []Temporal
}

// TODO: will need a way to apply a partial ordering to Temporal Nodes