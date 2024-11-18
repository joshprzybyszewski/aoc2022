package util

import (
	"fmt"
	"strconv"

	"github.com/joshprzybyszewski/aoc2022/puzzles/day01"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day02"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day03"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day04"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day05"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day06"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day07"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day08"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day09"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day10"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day11"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day12"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day13"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day14"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day15"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day16"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day17"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day18"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day19"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day20"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day21"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day22"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day23"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day24"
	"github.com/joshprzybyszewski/aoc2022/puzzles/day25"
)

func Solvers(
	day int,
) (part1, part2 func(string) (string, error)) {
	p1, p2 := IntSolvers(day)
	if p1 != nil && p2 != nil {
		return wrapIntSolver(p1), wrapIntSolver(p2)
	}
	switch day {
	case 25:
		return day25.One, nil

	}
	return nil, nil
}

func IntSolvers(
	day int,
) (part1, part2 func(string) (int, error)) {
	switch day {
	case 1:
		return day01.One, day01.Two
	case 2:
		return day02.One, day02.Two
	case 3:
		return day03.One, day03.Two
	case 4:
		return day04.One, day04.Two
	case 5:
		return day05.One, day05.Two
	case 6:
		return day06.One, day06.Two
	case 7:
		return day07.One, day07.Two
	case 8:
		return day08.One, day08.Two
	case 9:
		return day09.One, day09.Two
	case 10:
		return day10.One, day10.Two
	case 11:
		return day11.One, day11.Two
	case 12:
		return day12.One, day12.Two
	case 13:
		return day13.One, day13.Two
	case 14:
		return day14.One, day14.Two
	case 15:
		return day15.One, day15.Two
	case 16:
		return day16.One, day16.Two
	case 17:
		return day17.One, day17.Two
	case 18:
		return day18.One, day18.Two
	case 19:
		return day19.One, day19.Two
	case 20:
		return day20.One, day20.Two
	case 21:
		return day21.One, day21.Two
	case 22:
		return day22.One, day22.Two
	case 23:
		return day23.One, day23.Two
	case 24:
		return day24.One, day24.Two
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

func wrapUint64Solver(
	is func(string) (uint64, error),
) func(string) (string, error) {
	return func(input string) (string, error) {
		i, err := is(input)
		if err != nil {
			return ``, err
		}
		return fmt.Sprintf("%d", i), nil
	}
}
