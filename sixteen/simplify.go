package sixteen

import "fmt"

func simplify(
	vs []*valve,
) (map[string]*valve, map[string]*node) {
	valvesByName := make(map[string]*valve, len(vs))
	nodesByName := make(map[string]*node, len(vs))

	for _, v := range vs {
		valvesByName[v.name] = v
		if v.flow == 0 {
			continue
		}
		nodesByName[v.name] = newNode(
			v.name,
			v.flow,
		)
	}
	fmt.Printf("nodesByName: %+v\n", nodesByName)
	fmt.Printf("valvesByName: %+v\n", valvesByName)

	for name, n := range nodesByName {
		n.edges = getEdges(name, valvesByName, nodesByName)
	}

	return valvesByName, nodesByName
}

func getEdges(
	start string,
	vs map[string]*valve,
	nodesByName map[string]*node,
) []*edge {

	distByName := make(map[string]int, len(vs))
	var output []*edge

	pending := make([]string, 0, len(vs))
	pending = append(pending, start)

	for len(pending) > 0 {
		name := pending[0]
		if n, ok := nodesByName[name]; ok {
			output = append(output, newEdge(n, distByName[name]))
		}

		dv := distByName[name] + 1
		for _, d := range vs[name].dests {
			if _, ok := distByName[d]; ok {
				// this destination has already been seen
				continue
			}
			distByName[d] = dv
			pending = append(pending, d)
		}
		pending = pending[1:]
	}

	return output
}
