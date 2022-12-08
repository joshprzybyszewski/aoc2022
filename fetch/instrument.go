package fetch

import (
	"fmt"
	"runtime"
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

		answer, err := solver(input)

		runtime.ReadMemStats(&ms2)

		fmt.Printf("\ttotal alloc: %d\n", ms2.TotalAlloc-ms.TotalAlloc)
		fmt.Printf("\theap alloc: %d\n", ms2.HeapAlloc-ms.HeapAlloc)

		return answer, err
	}
}
