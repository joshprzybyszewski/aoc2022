package fifteen

const (
	numReports = 33
)

type report struct {
	sx, sy int
	bx, by int

	// calculated beacon distance
	dist int
}

func newReport(
	sx, sy int,
	bx, by int,
) report {
	dx := sx - bx
	if dx < 0 {
		dx = -dx
	}
	dy := sy - by
	if dy < 0 {
		dy = -dy
	}

	return report{
		sx:   sx,
		sy:   sy,
		bx:   bx,
		by:   by,
		dist: dx + dy,
	}
}

func (r *report) getSeenInRow(y int) (int, int, bool) {
	ry := y - r.sy
	if ry < 0 {
		ry = -ry
	}

	if ry > r.dist {
		return 0, 0, false
	}

	return r.sx - r.dist + ry, r.sx + r.dist - ry, true
}
