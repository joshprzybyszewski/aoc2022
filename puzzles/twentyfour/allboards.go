package twentyfour

type boardState uint16

const (
	// This should be 25 * 120, but we know that we never iterate
	// past a depth of 1024... so...
	numBoardStates boardState = 1024 // 25 rows, 120 columns
)

type allBoards struct {
	max boardState
	all [numBoardStates]board
}

func populatedAllBoards(
	initial *board,
) allBoards {
	ab := allBoards{}
	ab.all[0] = *initial
	return ab
}

func (ab *allBoards) getBoardAtState(
	bs boardState,
) *board {
	if bs > ab.max {
		for i := ab.max + 1; i <= bs; i++ {
			next(&ab.all[i-1], &ab.all[i])
		}
		ab.max = bs
	}
	return &ab.all[bs]
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
