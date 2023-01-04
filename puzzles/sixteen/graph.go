package sixteen

const (
	numNodes node = 15
)

type node uint8

type value uint8

type distance uint8

type graph struct {
	nodes [numNodes]value
	edges [numNodes][numNodes]distance

	startingPositions [numNodes]distance
}

func (g *graph) getDistance(
	s, d node,
) distance {
	return g.edges[s][d]
}

func (g *graph) getValue(
	n node,
) value {
	return g.nodes[n]
}
