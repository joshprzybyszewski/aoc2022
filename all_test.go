package main

import (
	"fmt"
	"testing"

	"github.com/joshprzybyszewski/aoc2022/util"
)

func BenchmarkAll(b *testing.B) {
	benchmarks := []struct {
		day     int
		answer1 string
		answer2 string
	}{{
		day:     1,
		answer1: `68787`,
		answer2: `198041`,
	}, {
		day:     2,
		answer1: `11873`,
		answer2: `12014`,
	}, {
		day:     3,
		answer1: `7863`,
		answer2: `2488`,
	}, {
		day:     4,
		answer1: `500`,
		answer2: `815`,
	}, {
		day:     5,
		answer1: `TWSGQHNHL`,
		answer2: `JNRSCDWPP`,
	}, {
		day:     6,
		answer1: `1300`,
		answer2: `3986`,
	}, {
		day:     7,
		answer1: `1915606`,
		answer2: `5025657`,
	}, {
		day:     8,
		answer1: `1809`,
		answer2: `479400`,
	}}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("Day %d", bm.day), func(b *testing.B) {
			input, err := util.Input(bm.day)
			if err != nil {
				b.Logf("Error fetching input: %q", err)
				b.Fail()
			}

			part1, part2 := util.Solvers(bm.day)

			b.Run(`Part One`, func(b *testing.B) {
				var answer string
				var err error
				for n := 0; n < b.N; n++ {
					answer, err = part1(input)
					if err != nil {
						b.Logf("got unexpected error: %q", err)
						b.Fail()
					}
					if answer != bm.answer1 {
						b.Logf("expected answer %q, but got %q", bm.answer1, answer)
						b.Fail()
					}
				}
			})

			b.Run(`Part Two`, func(b *testing.B) {
				var answer string
				var err error
				for n := 0; n < b.N; n++ {
					answer, err = part2(input)
					if err != nil {
						b.Logf("got unexpected error: %q", err)
						b.Fail()
					}
					if answer != bm.answer2 {
						b.Logf("expected answer %q, but got %q", bm.answer2, answer)
						b.Fail()
					}
				}
			})
		})
	}
}
