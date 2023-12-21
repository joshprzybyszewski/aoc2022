package eighteen

import (
	"slices"
)

func Two(
	input string,
) (int, error) {
	var l lagoon
	var i int
	for len(input) > 0 {
		l.edges[i], input = newEdge(input)
		i++
		input = input[1:]
	}
	l.numEdges = i

	gcfHor, gcfVer := convertEdges(&l)

	l.dig()

	// fmt.Printf("lagoon:\n%s\n", l.String())
	// fmt.Printf("gcfHor: %d\n", gcfHor)
	// fmt.Printf("gcfVer: %d\n", gcfVer)

	return l.numDug() * gcfHor * gcfVer, nil
}

func convertEdges(
	l *lagoon,
) (int, int) {
	for i := 0; i < l.numEdges; i++ {
		l.edges[i] = convertEdge(l.edges[i])
	}

	gcfHor, gcfVer := getGCF(l)

	for i := 0; i < l.numEdges; i++ {
		if l.edges[i].heading == east || l.edges[i].heading == west {
			l.edges[i].num /= gcfHor
		} else {
			l.edges[i].num /= gcfVer
		}

	}

	return gcfHor, gcfVer
}

func convertEdge(
	e edge,
) edge {
	switch e.color[6] {
	case '0':
		e.heading = east
	case '1':
		e.heading = south
	case '2':
		e.heading = west
	case '3':
		e.heading = north
	}

	e.num = 0
	for i := 1; i < 6; i++ {
		e.num *= 16 // TODO this could be <<= 8 probably
		e.num += int(e.color[i] - '0')
	}

	return e
}

func getGCF(
	l *lagoon,
) (int, int) {

	hors := make([]int, 0, l.numEdges)
	vers := make([]int, 0, l.numEdges)
	for i := 0; i < l.numEdges; i++ {
		if l.edges[i].heading == east || l.edges[i].heading == west {
			hors = append(hors, l.edges[i].num)
		} else {
			vers = append(vers, l.edges[i].num)
		}

	}
	slices.Sort(hors)
	slices.Sort(vers)

	gcfHor := 1
	n := 2
	for n < hors[0]/2 {
		if !isMultiple(hors, n) {
			n++
		}

		gcfHor *= n
		for i := range hors {
			hors[i] /= n
		}
	}

	gcfVer := 1
	n = 2
	for n < vers[0]/2 {
		if !isMultiple(vers, n) {
			n++
		}

		gcfVer *= n
		for i := range vers {
			vers[i] /= n
		}
	}

	return gcfHor, gcfVer
}

func isMultiple(
	vals []int,
	n int,
) bool {
	for _, v := range vals {
		if v%n != 0 {
			return false
		}
	}
	return true
}
