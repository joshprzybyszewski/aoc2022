package twentytwo

import (
	"strconv"
)

func convertInput(
	input string,
) (*space, []direction) {

	var r, c uint
	var r0, recent *space
	var r0s []*space
	checkR0 := func(s *space) {
		if r0 != nil {
			return
		}
		r0 = s
		r0s = append(r0s, r0)
	}
	closeRow := func() {
		r0.left = recent
		recent.right = r0
		r0 = nil
		recent = nil

		r++
		c = 1
	}

	getAbove := func(
		curR, curC uint,
	) *space {
		if len(r0s) == 0 || curR == 1 {
			return nil
		}
		var above *space
		for i := curR - 2; ; {
			above = getColumn(r0s[i], curC)
			if above != nil {
				return above
			}
			if i == 0 {
				break
			}
			i--
		}
		return nil
	}

	connectBottomRows := func() {
		var e, top *space
		for i := len(r0s) - 1; i > 0; i-- {
			// if this row has a higher max col that the previous,
			// then we need to fill something in

			for e = r0s[i].left; e.down == nil; e = e.left {
				for top = e; top.up != nil; top = top.up {

				}
				e.down = top
				top.up = e
			}
		}
	}

	var tmp *space
	r, c = 1, 1

	for i, ch := range input {
		switch ch {
		case ' ':
			c++
		case '.':
			tmp = newSpace(
				r, c,
				recent,
				getAbove(r, c),
			)
			checkR0(tmp)
			recent = tmp
			c++
		case '#':
			tmp = newWall(
				r, c,
				recent,
				getAbove(r, c),
			)
			checkR0(tmp)
			recent = tmp
			c++
		case '\n':
			if r0 == nil {
				// we've already seen a newline, therefore we're entering the directions phase.
				connectBottomRows()
				directions := getDirections(input, i+1)
				return r0s[0], directions
			}
			closeRow()
		}
	}

	panic(`should have found directions`)
}

func getDirections(
	input string,
	start int,
) []direction {

	output := make([]direction, 0, 128)

	h := right
	d0 := start
	var d int
	var err error
	for i := start; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			continue
		}

		d, err = strconv.Atoi(input[d0:i])
		if err != nil {
			panic(err)
		}
		d0 = i + 1
		output = append(output, direction{
			dist:    uint(d),
			heading: h,
		})

		switch input[i] {
		case 'L':
			if h == 0 {
				h = 3
			} else {
				h--
			}
		case 'R':
			if h == 3 {
				h = 0
			} else {
				h++
			}
		}
	}

	return output
}
