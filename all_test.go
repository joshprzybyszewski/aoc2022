package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/joshprzybyszewski/aoc2022/util"
)

func BenchmarkAll(b *testing.B) {
	now := time.Now()
	today := now.Day()
	if now.Year() > 2022 || today > 25 {
		today = 25
	}

	for day := 1; day <= today; day++ {
		b.Run(fmt.Sprintf("Day %d", day), func(b *testing.B) {
			input, err := util.Input(day)
			if err != nil {
				b.Logf("Error fetching input: %q", err)
				b.Skip()
			}

			answer1, err := util.Part1Answer(day)
			if err != nil {
				b.Logf("Error getting answers: %q", err)
				b.Skip()
			}

			part1, part2 := util.IntSolvers(day)
			part1String, part2String := util.Solvers(day)

			if part1 != nil {
				answer1, err := strconv.Atoi(answer1)
				if err != nil {
					b.Logf("got unexpected error: %q", err)
					b.Fail()
				}
				b.Run(`Part One`, func(b *testing.B) {
					var answer int
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
			} else {
				b.Run(`Part One`, func(b *testing.B) {
					var answer string
					var err error
					for n := 0; n < b.N; n++ {
						answer, err = part1String(input)
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
			}

			answer2, err := util.Part2Answer(day)
			if err != nil {
				b.Logf("Error getting answers: %q", err)
				return
			}

			if part2 != nil {
				answer2, err := strconv.Atoi(answer2)
				if err != nil {
					b.Logf("got unexpected error: %q", err)
					b.Fail()
				}
				b.Run(`Part Two`, func(b *testing.B) {
					var answer int
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
			} else {
				b.Run(`Part Two`, func(b *testing.B) {
					var answer string
					var err error
					for n := 0; n < b.N; n++ {
						answer, err = part2String(input)
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
			}
		})
	}
}
