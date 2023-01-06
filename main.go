package main

import (
	"flag"
	"fmt"

	"github.com/joshprzybyszewski/aoc2022/util"
)

var (
	shouldProfile = flag.Bool("profile", false, "if set, will produce a profile output")
	targetDay     = flag.Int("day", 0, "if set, will only run the given day")
)

func main() {
	flag.Parse()

	if *shouldProfile {
		defer util.Profile()()
	}

	if targetDay != nil && *targetDay > 0 && *targetDay <= 25 {
		err := runDay(*targetDay)
		if err != nil {
			panic(err)
		}
		return
	}

	for day := 1; day <= 25; day++ {
		err := runDay(day)
		if err != nil {
			panic(err)
		}
	}
}

func runDay(
	day int,
) error {
	input, err := util.Input(day)
	if err != nil {
		return err
	}

	part1, part2 := util.Solvers(day)
	part1, part2 = util.Instrument(part1, part2)
	part1, part2 = util.Submit(day, part1, part2)

	runParts(
		day,
		input,
		part1, part2,
	)
	return nil
}

func runParts(
	day int,
	input string,
	part1, part2 func(string) (string, error),
) {
	fmt.Printf("=====================================\n")
	fmt.Printf("Day %d\n", day)
	fmt.Printf("-------------------------------------\n")
	fmt.Printf("Part 1\n")
	answer, err := part1(input)
	if err != nil {
		fmt.Printf("Part 1 error: %q\n", err)
		panic(err)
	}
	fmt.Printf("Part 1 Answer: %q\n", answer)

	fmt.Printf("-------------------------------------\n")
	fmt.Printf("Part 2\n")
	answer, err = part2(input)
	if err != nil {
		fmt.Printf("Part 2 error: %q\n", err)
		panic(err)
	}
	fmt.Printf("Part 2 Answer: %q\n\n\n", answer)
}
