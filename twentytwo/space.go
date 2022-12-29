package twentytwo

func getColumn(
	r *space,
	c uint,
) *space {
	if r.col == c {
		return r
	}

	for r.right != nil && r.col < r.right.col {
		if r.right.col == c {
			return r.right
		}
		r = r.right
	}

	return nil
}

func move(
	s *space,
	d direction,
) *space {
	var next *space
	switch d.heading {
	case right:
		for i := uint(0); i < d.dist; i++ {
			next = s.right
			if next.isWall {
				return s
			}
			s = next
		}
	case left:
		for i := uint(0); i < d.dist; i++ {
			next = s.left
			if next.isWall {
				return s
			}
			s = next
		}
	case up:
		for i := uint(0); i < d.dist; i++ {
			next = s.up
			if next.isWall {
				return s
			}
			s = next
		}
	case down:
		for i := uint(0); i < d.dist; i++ {
			next = s.down
			if next.isWall {
				return s
			}
			s = next
		}
	}
	return s
}

type space struct {
	row uint
	col uint

	isWall bool

	up    *space
	right *space
	down  *space
	left  *space
}

func newSpace(
	r, c uint,
	left, up *space,
) *space {
	s := &space{
		row:    r,
		col:    c,
		isWall: false,
		up:     up,
		left:   left,
	}

	if left != nil {
		left.right = s
	}

	if up != nil {
		up.down = s
	}
	return s
}

func newWall(
	r, c uint,
	left, up *space,
) *space {
	s := &space{
		row:    r,
		col:    c,
		isWall: true,
		up:     up,
		left:   left,
	}

	if left != nil {
		left.right = s
	}
	if up != nil {
		up.down = s
	}

	return s
}
