package twentyfour

type position struct {
	row int
	col int
}

type path struct {
	cur      position
	bs       boardState
	numSteps int
}

func navigate(
	ab *allBoards,
	start position,
	startingState boardState,
	goal position,
) path {

	var handled [numBoardStates]board
	pending := make([]path, 0, 1028)
	pending = append(pending, path{
		cur: start,
		bs:  startingState,
	})

	var p path
	var bs boardState
	var b *board
	for len(pending) > 0 {
		p = pending[0]
		if p.cur == goal {
			return p
		}
		if handled[p.bs][p.cur.row][p.cur.col] != empty {
			pending = pending[1:]
			continue
		}

		handled[p.bs][p.cur.row][p.cur.col] = wall

		// check possible next positions for the updated board state
		bs = (p.bs + 1) % numBoardStates
		b = ab.getBoardAtState(bs)

		if b[p.cur.row][p.cur.col+1] == empty {
			pending = append(pending, path{
				cur: position{
					row: p.cur.row,
					col: p.cur.col + 1,
				},
				bs:       bs,
				numSteps: p.numSteps + 1,
			})
		}

		if b[p.cur.row+1][p.cur.col] == empty {
			pending = append(pending, path{
				cur: position{
					row: p.cur.row + 1,
					col: p.cur.col,
				},
				bs:       bs,
				numSteps: p.numSteps + 1,
			})
		}

		if b[p.cur.row][p.cur.col] == empty {
			pending = append(pending, path{
				cur: position{
					row: p.cur.row,
					col: p.cur.col,
				},
				bs:       bs,
				numSteps: p.numSteps + 1,
			})
		}

		if p.cur.row > 0 && b[p.cur.row-1][p.cur.col] == empty {
			pending = append(pending, path{
				cur: position{
					row: p.cur.row - 1,
					col: p.cur.col,
				},
				bs:       bs,
				numSteps: p.numSteps + 1,
			})
		}

		if b[p.cur.row][p.cur.col-1] == empty {
			pending = append(pending, path{
				cur: position{
					row: p.cur.row,
					col: p.cur.col - 1,
				},
				bs:       bs,
				numSteps: p.numSteps + 1,
			})
		}
		pending = pending[1:]
	}

	// goal not found
	return path{
		numSteps: -1,
	}
}
