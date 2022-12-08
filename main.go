package main

import (
	"fmt"
	"runtime"

	"github.com/joshprzybyszewski/aoc2022/fetch"
)

func main() {
	day := 8

	input, err := fetch.Input(day)
	if err != nil {
		panic(err)
	}

	part1, part2 := fetch.Solvers(day)

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
