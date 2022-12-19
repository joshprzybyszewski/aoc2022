package sixteen

const (
	// numNodes = 15 // real input has 15
	numNodes = 6 // test input has 6
)

// TODO can be uint8
type node int

// TODO PROBABLY can be uint8
type value int

// TODO can be uint8
type distance int

type graph struct {
	nodes [numNodes]value
	edges [numNodes][numNodes]distance

	startingPositions [numNodes]distance
}

func (g graph) getDistance(
	s, d node,
) distance {
	return g.edges[s][d]
}

func (g graph) getValue(
	n node,
) value {
	return g.nodes[n]
}
