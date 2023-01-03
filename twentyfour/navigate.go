package twentyfour

type position struct {
	row int
	col int
}

type boardState uint16

type path struct {
	cur  position
	bs   boardState
	prev *path
}

func navigate(
	ab *allBoards,
	start position,
	startingState boardState,
	goal position,
) *path {
	var handled allBoards
	n := navigator{
		ab:      ab,
		pending: make([]*path, 0, 1028),
		goal:    goal,
		handled: &handled,
	}
	n.pending = append(n.pending, &path{
		cur: start,
		bs:  startingState,
	})

	return n.solve()
}

type navigator struct {
	ab      *allBoards
	pending []*path
	goal    position

	handled *allBoards
}

func (n *navigator) solve() *path {
	var p *path
	var bs boardState
	var b *board
	for len(n.pending) > 0 {
		p = n.pending[0]
		if p.cur == n.goal {
			return p
		}
		if n.hasPreviouslyHandled(p) {
			n.pending = n.pending[1:]
			continue
		}

		// check possible next positions for the updated board state
		bs = (p.bs + 1) % numBoardStates
		b = n.ab.getBoardAtState(bs)

		if b[p.cur.row][p.cur.col+1] == empty {
			n.pending = append(n.pending, &path{
				cur: position{
					row: p.cur.row,
					col: p.cur.col + 1,
				},
				bs:   bs,
				prev: p,
			})
		}

		if b[p.cur.row+1][p.cur.col] == empty {
			n.pending = append(n.pending, &path{
				cur: position{
					row: p.cur.row + 1,
					col: p.cur.col,
				},
				bs:   bs,
				prev: p,
			})
		}

		if b[p.cur.row][p.cur.col] == empty {
			n.pending = append(n.pending, &path{
				cur: position{
					row: p.cur.row,
					col: p.cur.col,
				},
				bs:   bs,
				prev: p,
			})
		}

		if p.cur.row > 0 && b[p.cur.row-1][p.cur.col] == empty {
			n.pending = append(n.pending, &path{
				cur: position{
					row: p.cur.row - 1,
					col: p.cur.col,
				},
				bs:   bs,
				prev: p,
			})
		}

		if b[p.cur.row][p.cur.col-1] == empty {
			n.pending = append(n.pending, &path{
				cur: position{
					row: p.cur.row,
					col: p.cur.col - 1,
				},
				bs:   bs,
				prev: p,
			})
		}
		n.pending = n.pending[1:]
	}
	return nil
}

func (n *navigator) hasPreviouslyHandled(p *path) bool {
	b := n.handled.getBoardAtState(p.bs)
	if b[p.cur.row][p.cur.col] != empty {
		return true
	}
	b[p.cur.row][p.cur.col] = wall
	return false
}
