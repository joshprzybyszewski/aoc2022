package twentyfour

import "fmt"

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
	goal position,
) *path {
	n := navigator{
		ab: ab,
		pending: []*path{{
			cur: start,
			bs:  0,
		}},
		goal: goal,
	}

	return n.solve()
}

type navigator struct {
	ab      *allBoards
	pending []*path
	goal    position

	handled allBoards
}

func (n *navigator) solve() *path {
	var p *path
	total := 0
	for len(n.pending) > 0 {
		total++
		p = n.pending[0]
		if p.cur == n.goal {
			return p
		}
		n.handle(p)
		n.pending = n.pending[1:]
	}
	fmt.Printf("handled %d\n", total)
	return nil
}

func (n *navigator) handle(p *path) {
	if n.hasPreviouslyHandled(p) {
		return
	}

	// check possible next positions for the updated board state
	bs := (p.bs + 1) % numBoardStates
	b := n.ab.getBoardAtState(bs)

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
}

func (n *navigator) hasPreviouslyHandled(p *path) bool {
	b := n.handled.getBoardAtState(p.bs)
	if b[p.cur.row][p.cur.col] != empty {
		return true
	}
	b[p.cur.row][p.cur.col] = wall
	return false
}
