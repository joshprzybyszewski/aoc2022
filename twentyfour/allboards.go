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
		ab[i] = next(ab[i-1])
	}
	return ab
}

func (ab *allBoards) getBoardAtState(
	i boardState,
) *board {
	return &ab[i]
}
