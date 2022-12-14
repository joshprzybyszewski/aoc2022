package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/joshprzybyszewski/aoc2022/util"
)

func BenchmarkAll(b *testing.B) {
	now := time.Now()
	today := now.Day()
	if now.Year() > 2022 || today > 25 {
		today = 25
	} else {
		today = 12
	}

	// for day := 7; day <= 7; day++ {
	for day := 1; day <= today; day++ {
		b.Run(fmt.Sprintf("Day %d", day), func(b *testing.B) {
			input, err := util.Input(day)
			if err != nil {
				b.Logf("Error fetching input: %q", err)
				b.Fail()
			}

			answer1, answer2, err := util.Answers(day)
			if err != nil {
				b.Logf("Error getting answers: %q", err)
				b.Fail()
			}

			part1, part2 := util.Solvers(day)

			b.Run(`Part One`, func(b *testing.B) {
				var answer string
				var err error
				for n := 0; n < b.N; n++ {
					answer, err = part1(input)
					if err != nil {
						b.Logf("got unexpected error: %q", err)
						b.Fail()
					}
					if answer != answer1 {
						b.Logf("expected answer %q, but got %q", answer1, answer)
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
					if answer != answer2 {
						b.Logf("expected answer %q, but got %q", answer2, answer)
						b.Fail()
					}
				}
			})
		})
	}
}
