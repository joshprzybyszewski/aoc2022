package seventeen

import (
	"fmt"
)

const (
	numRocksPart2 = 1000000000000
)

type reduction struct {
	jetIndex   int
	numReduced int
	rockIndex  int

	chamber chamber
}

func (r *reduction) String() string {
	s := fmt.Sprintf("\nReduced %4d rows\n", r.numReduced)
	s += fmt.Sprintf("\tHeight:        %3d\n", r.chamber.minEmptyRow)
	s += fmt.Sprintf("\tPending:         %d\n", r.chamber.pendingIndex)
	s += fmt.Sprintf("\tJet Index:   %5d\n", r.jetIndex)
	s += fmt.Sprintf("\tRock Index: %6d\n\n\n", r.rockIndex)
	return s
}

func (r *reduction) same(other *reduction) bool {
	if r.jetIndex != other.jetIndex {
		return false
	}
	if r.chamber.minEmptyRow != other.chamber.minEmptyRow {
		return false
	}
	if r.chamber.pendingIndex != other.chamber.pendingIndex {
		return false
	}

	for row := 0; row < r.chamber.minEmptyRow; row++ {
		if r.chamber.settled[row] != other.chamber.settled[row] {
			return false
		}
	}

	return true
}

func Two(
	input string,
) (int, error) {
	c := newChamber()
	i := 0

	numRowsReduced := 0
	var r int

	reductions := make([]*reduction, 0, 1024)
	var same *reduction

	for nr := 0; nr < numRocksPart2; nr++ {
		r = c.reduce()
		if r > 0 {
			numRowsReduced += r
			red := &reduction{
				jetIndex:   i,
				numReduced: numRowsReduced,
				rockIndex:  nr,
				chamber:    c,
			}
			same = nil
			for _, other := range reductions {
				if red.same(other) {
					same = other
					break
				}
			}
			if same == nil {
				reductions = append(reductions, red)
			} else {
				dRow := red.numReduced - same.numReduced
				dNR := red.rockIndex - same.rockIndex
				for nr+dNR < numRocksPart2 {
					numRowsReduced += dRow
					nr += dNR
				}
			}

		}

		for {
			switch input[i] {
			case '<':
				c.pushLeft()
			case '>':
				c.pushRight()
			default:
				panic(input[i])
			}

			i++
			if i == len(input)-1 {
				i = 0
			}

			if !c.fall() {
				break
			}
		}
	}

	return numRowsReduced + c.minEmptyRow, nil
}
