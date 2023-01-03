package twentyfour

const (
	numBoardStates boardState = 25 * 120 // 25 rows, 120 columns
	// but I think this reduces down to 5 rows, 24 columns ==> 120 states
)

type allBoards [numBoardStates]board

func newAllBoards(
	initial board,
) allBoards {
	ab := allBoards{}
	ab[0] = initial
	for i := 1; i < len(ab); i++ {
		next(&ab[i-1], &ab[i])
	}
	return ab
}

func next(
	src, dst *board,
) {
	var c int
	var s square
	for r := range src {
		for c, s = range src[r] {
			if s == wall {
				dst[r][c] = wall
				continue
			}

			if s&right == right {
				if c == len(src[r])-2 {
					dst[r][1] |= right
				} else {
					dst[r][c+1] |= right
				}
			}
			if s&down == down {
				if r == len(src)-2 {
					dst[1][c] |= down
				} else {
					dst[r+1][c] |= down
				}
			}
			if s&left == left {
				if c == 1 {
					dst[r][len(src[r])-2] |= left
				} else {
					dst[r][c-1] |= left
				}
			}
			if s&up == up {
				if r == 1 {
					dst[len(src)-2][c] |= up
				} else {
					dst[r-1][c] |= up
				}
			}
		}
	}
}

func (ab *allBoards) getBoardAtState(
	i boardState,
) *board {
	return &ab[i]
}
