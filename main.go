package main

import (
	"flag"
	"fmt"

	"github.com/joshprzybyszewski/aoc2022/util"
)

var (
	shouldProfile = flag.Bool("profile", false, "if set, will produce a profile output")
)

func main() {
	flag.Parse()

	if *shouldProfile {
		defer util.Profile()()
	}

	day := 18
	input, err := util.Input(day)
	if err != nil {
		panic(err)
	}

	part1, part2 := util.Solvers(day)
	part1, part2 = util.Instrument(part1, part2)
	part1, part2 = util.Submit(day, part1, part2)

	runParts(
		day,
		input,
		part1, part2,
	)
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
	fmt.Printf("Part 2 Answer: %q\n", answer)
}
