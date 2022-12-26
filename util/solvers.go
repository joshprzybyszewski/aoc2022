package util

import (
	"fmt"
	"strconv"

	"github.com/joshprzybyszewski/aoc2022/eight"
	"github.com/joshprzybyszewski/aoc2022/eighteen"
	"github.com/joshprzybyszewski/aoc2022/eleven"
	"github.com/joshprzybyszewski/aoc2022/fifteen"
	"github.com/joshprzybyszewski/aoc2022/five"
	"github.com/joshprzybyszewski/aoc2022/four"
	"github.com/joshprzybyszewski/aoc2022/fourteen"
	"github.com/joshprzybyszewski/aoc2022/nine"
	"github.com/joshprzybyszewski/aoc2022/nineteen"
	"github.com/joshprzybyszewski/aoc2022/one"
	"github.com/joshprzybyszewski/aoc2022/seven"
	"github.com/joshprzybyszewski/aoc2022/seventeen"
	"github.com/joshprzybyszewski/aoc2022/six"
	"github.com/joshprzybyszewski/aoc2022/sixteen"
	"github.com/joshprzybyszewski/aoc2022/ten"
	"github.com/joshprzybyszewski/aoc2022/thirteen"
	"github.com/joshprzybyszewski/aoc2022/three"
	"github.com/joshprzybyszewski/aoc2022/twelve"
	"github.com/joshprzybyszewski/aoc2022/twenty"
	"github.com/joshprzybyszewski/aoc2022/twentyone"
	"github.com/joshprzybyszewski/aoc2022/two"
)

func Solvers(
	day int,
) (part1, part2 func(string) (string, error)) {
	p1, p2 := IntSolvers(day)
	if p1 != nil && p2 != nil {
		return wrapIntSolver(p1), wrapIntSolver(p2)
	}
	switch day {
	case 5:
		return five.One, five.Two
	case 10:
		return wrapIntSolver(ten.One), ten.Two
	case 11:
		return wrapIntSolver(eleven.One), wrapInt64Solver(eleven.Two)
	}
	return nil, nil
}

func IntSolvers(
	day int,
) (part1, part2 func(string) (int, error)) {
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
		// the answers are strings
		return nil, nil
		// return five.One, five.Two
	case 6:
		return six.One, six.Two
	case 7:
		return seven.One, seven.Two
	case 8:
		return eight.One, eight.Two
	case 9:
		return nine.One, nine.Two
	case 10:
		return ten.One, nil
	case 11:
		return eleven.One, nil
	case 12:
		return twelve.One, twelve.Two
	case 13:
		return thirteen.One, thirteen.Two
	case 14:
		return fourteen.One, fourteen.Two
	case 15:
		return fifteen.One, fifteen.Two
	case 16:
		return sixteen.One, sixteen.Two
	case 17:
		return seventeen.One, seventeen.Two
	case 18:
		return eighteen.One, eighteen.Two
	case 19:
		return nineteen.One, nineteen.Two
	case 20:
		return twenty.One, twenty.Two
	case 21:
		return twentyone.One, twentyone.Two
	}
	return nil, nil
}

func wrapIntSolver(
	is func(string) (int, error),
) func(string) (string, error) {
	return func(input string) (string, error) {
		i, err := is(input)
		if err != nil {
			return ``, err
		}
		return strconv.Itoa(i), nil
	}
}

func wrapInt64Solver(
	is func(string) (int64, error),
) func(string) (string, error) {
	return func(input string) (string, error) {
		i, err := is(input)
		if err != nil {
			return ``, err
		}
		return fmt.Sprintf("%d", i), nil
	}
}
