package sixteen

type node struct {
	name  string
	value int
	edges []*edge
}

func newNode(
	name string,
	value int,
) *node {
	return &node{
		name:  name,
		value: value,
	}
}
