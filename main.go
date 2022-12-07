package main

import (
	"fmt"
	"runtime"

	"github.com/joshprzybyszewski/aoc2022/eight"
	"github.com/joshprzybyszewski/aoc2022/fetch"
	"github.com/joshprzybyszewski/aoc2022/five"
	"github.com/joshprzybyszewski/aoc2022/four"
	"github.com/joshprzybyszewski/aoc2022/one"
	"github.com/joshprzybyszewski/aoc2022/seven"
	"github.com/joshprzybyszewski/aoc2022/six"
	"github.com/joshprzybyszewski/aoc2022/three"
	"github.com/joshprzybyszewski/aoc2022/two"
)

func main() {
	day := 6

	input, err := fetch.Input(day)
	if err != nil {
		panic(err)
	}

	var part1, part2 func(string) (string, error)
	switch day {
	case 1:
		part1 = one.One
		part2 = one.Two
	case 2:
		part1 = two.One
		part2 = two.Two
	case 3:
		part1 = three.One
		part2 = three.Two
	case 4:
		part1 = four.One
		part2 = four.Two
	case 5:
		part1 = five.One
		part2 = five.Two
	case 6:
		part1 = six.One
		part2 = six.Two
	case 7:
		part1 = seven.One
		part2 = seven.Two
	case 8:
		part1 = eight.One
		part2 = eight.Two
	}

	err = runParts(
		day,
		input,
		part1, part2,
	)
	if err != nil {
		panic(err)
	}
}

func runParts(
	day int,
	input string,
	part1, part2 func(string) (string, error),
) error {
	fmt.Printf("Day %d\n", day)
	ms := &runtime.MemStats{}
	runtime.ReadMemStats(ms)
	answer, err := part1(input)
	if err != nil {
		panic(err)
	}
	ms2 := &runtime.MemStats{}
	runtime.ReadMemStats(ms2)

	fmt.Printf("total alloc: %d\n", ms2.TotalAlloc-ms.TotalAlloc)
	fmt.Printf("heap alloc: %d\n", ms2.HeapAlloc-ms.HeapAlloc)

	var resp string
	fmt.Printf("Submit part 1 answer? (Y/n) %q\n", answer)
	fmt.Scanf("%s", &resp)
	if len(resp) > 0 && (resp == `y` || resp == `Y`) {
		resp, err := fetch.Answer(day, 1, answer)
		if err != nil {
			fmt.Printf("error while submitting: %v\n", err)
		} else {
			fmt.Printf("Successfully submitted: %q\n", resp)
		}
	}

	runtime.ReadMemStats(ms)

	answer, err = part2(input)
	if err != nil {
		panic(err)
	}

	runtime.ReadMemStats(ms2)
	fmt.Printf("total alloc: %d\n", ms2.TotalAlloc-ms.TotalAlloc)
	fmt.Printf("heap alloc: %d\n", ms2.HeapAlloc-ms.HeapAlloc)

	fmt.Printf("Submit part 2 answer? (Y/n) %q\n", answer)
	fmt.Scanf("%s", &resp)
	if len(resp) > 0 && (resp == `y` || resp == `Y`) {
		resp, err := fetch.Answer(day, 2, answer)
		if err != nil {
			fmt.Printf("error while submitting: %v\n", err)
		} else {
			fmt.Printf("Successfully submitted: %q\n", resp)
		}
	}
	return nil
}
