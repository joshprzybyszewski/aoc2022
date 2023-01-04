package sixteen

import (
	"sort"
)

func buildGraph(
	startingNode string,
	valves *valves,
) graph {
	nameToIndex := make(map[string]int, numNodes)
	vs := make(map[string]*valve, len(valves))

	indexesWithFlow := make([]int, 0, numNodes)

	for i, v := range valves {
		vs[v.name] = &valves[i]
		if v.flow > 0 {
			indexesWithFlow = append(indexesWithFlow, i)
		}
	}

	sort.Slice(indexesWithFlow, func(i, j int) bool {
		return valves[indexesWithFlow[i]].flow > valves[indexesWithFlow[j]].flow
	})

	nodes := [numNodes]value{}
	for i, vi := range indexesWithFlow {
		nodes[i] = value(valves[vi].flow)
		nameToIndex[valves[vi].name] = i
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
