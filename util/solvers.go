package util

import (
	"fmt"
	"strconv"

	"github.com/joshprzybyszewski/aoc2022/puzzles/eight"
	"github.com/joshprzybyszewski/aoc2022/puzzles/eighteen"
	"github.com/joshprzybyszewski/aoc2022/puzzles/eleven"
	"github.com/joshprzybyszewski/aoc2022/puzzles/fifteen"
	"github.com/joshprzybyszewski/aoc2022/puzzles/five"
	"github.com/joshprzybyszewski/aoc2022/puzzles/four"
	"github.com/joshprzybyszewski/aoc2022/puzzles/fourteen"
	"github.com/joshprzybyszewski/aoc2022/puzzles/nine"
	"github.com/joshprzybyszewski/aoc2022/puzzles/nineteen"
	"github.com/joshprzybyszewski/aoc2022/puzzles/one"
	"github.com/joshprzybyszewski/aoc2022/puzzles/seven"
	"github.com/joshprzybyszewski/aoc2022/puzzles/seventeen"
	"github.com/joshprzybyszewski/aoc2022/puzzles/six"
	"github.com/joshprzybyszewski/aoc2022/puzzles/sixteen"
	"github.com/joshprzybyszewski/aoc2022/puzzles/ten"
	"github.com/joshprzybyszewski/aoc2022/puzzles/thirteen"
	"github.com/joshprzybyszewski/aoc2022/puzzles/three"
	"github.com/joshprzybyszewski/aoc2022/puzzles/twelve"
	"github.com/joshprzybyszewski/aoc2022/puzzles/twenty"
	"github.com/joshprzybyszewski/aoc2022/puzzles/twentyfive"
	"github.com/joshprzybyszewski/aoc2022/puzzles/twentyfour"
	"github.com/joshprzybyszewski/aoc2022/puzzles/twentyone"
	"github.com/joshprzybyszewski/aoc2022/puzzles/twentythree"
	"github.com/joshprzybyszewski/aoc2022/puzzles/twentytwo"
	"github.com/joshprzybyszewski/aoc2022/puzzles/two"
)

func Solvers(
	day int,
) (part1, part2 func(string) (string, error)) {
	p1, p2 := IntSolvers(day)
	if p1 != nil && p2 != nil {
		return wrapIntSolver(p1), wrapIntSolver(p2)
	}
	switch day {
	case 6:
		return wrapInt64Solver(six.One), wrapIntSolver(six.Two)
	case 10:
		return wrapIntSolver(ten.One), ten.Two
	case 11:
		return wrapIntSolver(eleven.One), wrapInt64Solver(eleven.Two)
	case 25:
		return twentyfive.One, nil

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
		return five.One, five.Two
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
	case 22:
		return twentytwo.One, twentytwo.Two
	case 23:
		return twentythree.One, twentythree.Two
	case 24:
		return twentyfour.One, twentyfour.Two
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
