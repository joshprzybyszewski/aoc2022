package sixteen

type edge struct {
	dest   *node
	weight int
}

func newEdge(
	d *node,
	v int,
) *edge {
	return &edge{
		dest:   d,
		weight: v,
	}
}
