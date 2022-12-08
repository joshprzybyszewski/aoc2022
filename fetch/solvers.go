package fetch

import (
	"github.com/joshprzybyszewski/aoc2022/eight"
	"github.com/joshprzybyszewski/aoc2022/five"
	"github.com/joshprzybyszewski/aoc2022/four"
	"github.com/joshprzybyszewski/aoc2022/one"
	"github.com/joshprzybyszewski/aoc2022/seven"
	"github.com/joshprzybyszewski/aoc2022/six"
	"github.com/joshprzybyszewski/aoc2022/three"
	"github.com/joshprzybyszewski/aoc2022/two"
)

func Solvers(
	day int,
) (part1, part2 func(string) (string, error)) {
	switch day {
	case 1:
		return one.One, one.Two
	case 2:
		return two.One, two.Two
	case 3:
		return three.One, three.Two
	case 4:
		return four.One, four.Two
	case 5:
		return five.One, five.Two
	case 6:
		return six.One, six.Two
	case 7:
		return seven.One, seven.Two
	case 8:
		return eight.One, eight.Two
	}
	return nil, nil
}
