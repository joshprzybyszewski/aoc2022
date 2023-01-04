package twentythree

import "sort"

type coordRange struct {
	x0, x1 int
	y0, y1 int
}

// I want to use a k-d tree for this, where k == 2
type kdTree struct {
	root *xsplit
}

func newKDTree(cs []coord) *kdTree {
	return &kdTree{
		root: newXSplit(cs),
	}
}

func (t *kdTree) search(r coordRange) []coord {
	if t == nil {
		return nil
	}
	return t.root.search(&r)
}

type xsplit struct {
	c coord

	left  *ysplit
	right *ysplit
}

func newXSplit(
	cs []coord,
) *xsplit {
	if len(cs) == 0 {
		return nil
	}
	if len(cs) == 1 {
		return &xsplit{
			c: cs[0],
		}
	}

	sort.Slice(cs, func(i, j int) bool {
		if cs[i].x == cs[j].x {
			return cs[i].y < cs[j].y
		}
		return cs[i].x < cs[j].x
	})

	m := len(cs) / 2

	return &xsplit{
		c:     cs[m],
		left:  newYSplit(cs[:m]),
		right: newYSplit(cs[m+1:]),
	}
}

func (n *xsplit) all() []coord {
	if n == nil {
		return nil
	}
	output := n.left.all()
	output = append(output, n.c)
	output = append(output, n.right.all()...)
	return output
}

func (n *xsplit) search(
	r *coordRange,
) []coord {
	var output []coord
	if r.x0 <= n.c.x && n.c.x <= r.x1 &&
		r.y0 <= n.c.y && n.c.y <= r.y1 {
		output = append(output, n.c)
	}

	if n.left != nil && r.x0 <= n.c.x {
		output = append(output, n.left.search(r)...)

	}

	if n.right != nil && r.x1 >= n.c.x {
		output = append(output, n.right.search(r)...)
	}

	return output
}

type ysplit struct {
	c coord

	up   *xsplit
	down *xsplit
}

func newYSplit(
	cs []coord,
) *ysplit {
	if len(cs) == 0 {
		return nil
	}
	if len(cs) == 1 {
		return &ysplit{
			c: cs[0],
		}
	}

	sort.Slice(cs, func(i, j int) bool {
		if cs[i].y == cs[j].y {
			return cs[i].x < cs[j].x
		}
		return cs[i].y < cs[j].y
	})

	m := len(cs) / 2

	return &ysplit{
		c:    cs[m],
		up:   newXSplit(cs[:m]),
		down: newXSplit(cs[m+1:]),
	}
}

func (n *ysplit) all() []coord {
	if n == nil {
		return nil
	}
	output := n.up.all()
	output = append(output, n.c)
	output = append(output, n.down.all()...)
	return output
}

func (n *ysplit) search(
	r *coordRange,
) []coord {
	var output []coord
	if r.x0 <= n.c.x && n.c.x <= r.x1 &&
		r.y0 <= n.c.y && n.c.y <= r.y1 {
		output = append(output, n.c)
	}

	if n.up != nil && r.y0 <= n.c.y {
		output = append(output, n.up.search(r)...)

	}

	if n.down != nil && r.y1 >= n.c.y {
		output = append(output, n.down.search(r)...)
	}

	return output
}
