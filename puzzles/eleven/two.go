package eleven

const (
	oneMillion = 999999
)

func (u *universe) olderShortestPath(
	i, j int,
) int64 {
	start := u.universes[i]
	end := u.universes[j]
	if end.row < start.row { // TODO This will never be true
		end.row, start.row = start.row, end.row
	}
	if end.col < start.col {
		end.col, start.col = start.col, end.col
	}

	endRow := uint64(end.row)
	endCol := uint64(end.col)

	numExpanded := uint64(0)
	tmp := 0

	for tmp = start.row + 1; tmp < end.row; tmp++ {
		if !u.rowsWith[tmp] {
			numExpanded += oneMillion
		}
	}

	endRow += numExpanded
	numExpanded = 0

	for tmp = start.col + 1; tmp < end.col; tmp++ {
		if !u.colsWith[tmp] {
			numExpanded += oneMillion
		}
	}

	endCol += numExpanded

	return int64((endCol - uint64(start.col)) + (endRow - uint64(start.row)))
}

func Two(
	input string,
) (int64, error) {
	u := newUniverse(input)

	var total int64

	for i := 0; i < len(u.universes); i++ {
		for j := i + 1; j < len(u.universes); j++ {
			total += int64(u.olderShortestPath(i, j))
		}
	}

	return total, nil
}
