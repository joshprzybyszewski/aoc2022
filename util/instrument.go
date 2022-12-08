package util

import (
	"fmt"
	"runtime"
	"time"
)

func Instrument(
	part1, part2 func(string) (string, error),
) (_, __ func(string) (string, error)) {
	return report(part1), report(part2)
}

func report(solver func(string) (string, error)) func(string) (string, error) {
	ms := runtime.MemStats{}
	ms2 := runtime.MemStats{}

	return func(input string) (string, error) {
		runtime.ReadMemStats(&ms)

		t0 := time.Now()

		answer, err := solver(input)

		t1 := time.Now()
		runtime.ReadMemStats(&ms2)

		fmt.Printf("\tduration  :\t%s\n", t1.Sub(t0))
		fmt.Printf("\theap alloc:\t%d\n", ms2.HeapAlloc-ms.HeapAlloc)

		return answer, err
	}
}
