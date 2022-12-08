package main

import (
	"fmt"
	"testing"

	"github.com/joshprzybyszewski/aoc2022/util"
)

func BenchmarkAll(b *testing.B) {
	/*
		goos: linux
		goarch: amd64
		pkg: github.com/joshprzybyszewski/aoc2022
		cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz
		BenchmarkAll/Day_1/Part_One-4         	   17919	     67104 ns/op	   40983 B/op	       3 allocs/op
		BenchmarkAll/Day_1/Part_Two-4         	   14316	     85408 ns/op	   45096 B/op	      13 allocs/op
		BenchmarkAll/Day_2/Part_One-4         	    3313	    390453 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_2/Part_Two-4         	    2690	    387169 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_3/Part_One-4         	   21720	     55214 ns/op	    4868 B/op	       2 allocs/op
		BenchmarkAll/Day_3/Part_Two-4         	   20006	     59070 ns/op	   13780 B/op	      55 allocs/op
		BenchmarkAll/Day_4/Part_One-4         	    2719	    433709 ns/op	  112387 B/op	    3002 allocs/op
		BenchmarkAll/Day_4/Part_Two-4         	     823	   1477118 ns/op	   87342 B/op	    5460 allocs/op
		BenchmarkAll/Day_5/Part_One-4         	    1386	    915075 ns/op	   66394 B/op	    2615 allocs/op
		BenchmarkAll/Day_5/Part_Two-4         	    1408	    904026 ns/op	   66426 B/op	    2614 allocs/op
		BenchmarkAll/Day_6/Part_One-4         	  198666	      5837 ns/op	       4 B/op	       1 allocs/op
		BenchmarkAll/Day_6/Part_Two-4         	  235473	      4558 ns/op	       4 B/op	       1 allocs/op
		BenchmarkAll/Day_7/Part_One-4         	    4672	    248655 ns/op	   79861 B/op	    1811 allocs/op
		BenchmarkAll/Day_7/Part_Two-4         	    6434	    242365 ns/op	   76820 B/op	    1705 allocs/op
		BenchmarkAll/Day_8/Part_One-4         	    8934	    135256 ns/op	  106964 B/op	     202 allocs/op
		BenchmarkAll/Day_8/Part_Two-4         	    2401	    434516 ns/op	   93192 B/op	     102 allocs/op
	*/

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
