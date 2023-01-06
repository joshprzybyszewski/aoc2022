package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"

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

	err := runAll()
	if err != nil {
		panic(err)
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

	fmt.Printf("=====================================\n")
	fmt.Printf("Day %d\n", day)
	fmt.Printf("-------------------------------------\n")
	fmt.Printf("Part 1\n")
	answer, err := part1(input)
	if err != nil {
		fmt.Printf("Part 1 error: %q\n", err)
		return err
	}
	fmt.Printf("Part 1 Answer: %q\n", answer)

	fmt.Printf("-------------------------------------\n")
	fmt.Printf("Part 2\n")
	answer, err = part2(input)
	if err != nil {
		fmt.Printf("Part 2 error: %q\n", err)
		return err
	}
	fmt.Printf("Part 2 Answer: %q\n\n\n", answer)

	return nil
}

func runAll() error {
	var inputs [25]string
	var solvers [25][2]func(string) (string, error)
	var answers [25][2]string

	for day := 1; day <= 25; day++ {
		input, err := util.Input(day)
		if err != nil {
			return err
		}
		inputs[day-1] = input

		part1, part2 := util.Solvers(day)
		solvers[day-1][0] = part1
		solvers[day-1][1] = part2
	}

	getAnswer := func(day, part int) {
		if answers[day-1][part] != `` {
			return
		}
		answers[day-1][part], _ = solvers[day-1][part](inputs[day-1])
	}

	// these solvers already use concurrency: don't tromp on them.
	getAnswer(11, 1)
	getAnswer(19, 0)
	getAnswer(19, 1)
	getAnswer(16, 0)
	getAnswer(16, 1)
	getAnswer(23, 0)
	getAnswer(23, 1)

	var wg sync.WaitGroup
	work := make(chan int, 50)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for w := range work {
				getAnswer(w/2+1, w%2)
				wg.Done()
			}
		}()
	}

	for i := 0; i < 49; i++ {
		wg.Add(1)
		work <- i
	}

	wg.Wait()

	for day := 1; day <= 25; day++ {
		fmt.Printf("=====================\n")
		fmt.Printf("Day %2d\n", day)
		fmt.Printf("\tPart 1: %q\n", answers[day-1][0])
		fmt.Printf("\tPart 2: %q\n\n", answers[day-1][1])
	}

	return nil
}
