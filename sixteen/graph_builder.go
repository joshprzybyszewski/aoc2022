package sixteen

func buildGraph(
	startingNode string,
	valves []*valve,
) graph {
	nameToIndex := make(map[string]int, numNodes)
	var names [numNodes]string
	vs := make(map[string]*valve, len(valves))

	nodes := [numNodes]value{}
	ni := 0
	for _, v := range valves {
		vs[v.name] = v
		if v.flow > 0 {
			nodes[ni] = value(v.flow)
			nameToIndex[v.name] = ni
			names[ni] = v.name
			ni++
		}
	}

	edges := [numNodes][numNodes]distance{}

	for name, index := range nameToIndex {
		edges[index] = getEdges(
			name,
			vs,
			nameToIndex,
		)
	}

	startingPositions := getEdges(
		startingNode,
		vs,
		nameToIndex,
	)

	// for n := 0; n < numNodes; n++ {
	// 	fmt.Printf("Node %q has flow rate: %2d\n", names[n], nodes[n])
	// 	for d := 0; d < numNodes; d++ {
	// 		fmt.Printf("\tCan reach %q in %2d minutes\n", names[d], edges[n][d])
	// 	}
	// }

	// fmt.Printf("STARTING POSITIONS\n")
	// for n := 0; n < numNodes; n++ {
	// 	fmt.Printf("\tNode %q at time %2d\n", names[n], startingPositions[n])
	// }

	return graph{
		nodes:             nodes,
		edges:             edges,
		startingPositions: startingPositions,
	}
}

func getEdges(
	start string,
	vs map[string]*valve,
	nameToIndex map[string]int,
) [numNodes]distance {

	distByName := make(map[string]distance, len(vs))

	pending := make([]string, 0, len(vs))
	pending = append(pending, start)

	var name, dest string
	var dv distance
	var ok bool
	for len(pending) > 0 {
		name = pending[0]

		dv = distByName[name] + 1
		for _, dest = range vs[name].dests {
			if _, ok = distByName[dest]; ok {
				// this destination has already been seen
				continue
			}
			distByName[dest] = dv
			pending = append(pending, dest)
		}
		pending = pending[1:]
	}

	var output [numNodes]distance

	for name, index := range nameToIndex {
		output[index] = distByName[name]
	}

	return output
}