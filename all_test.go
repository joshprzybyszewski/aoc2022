package main

import (
	"testing"

	"github.com/joshprzybyszewski/aoc2022/fetch"
	"github.com/joshprzybyszewski/aoc2022/five"
	"github.com/joshprzybyszewski/aoc2022/four"
	"github.com/joshprzybyszewski/aoc2022/one"
	"github.com/joshprzybyszewski/aoc2022/seven"
	"github.com/joshprzybyszewski/aoc2022/six"
	"github.com/joshprzybyszewski/aoc2022/three"
	"github.com/joshprzybyszewski/aoc2022/two"
)

func BenchmarkAll(b *testing.B) {
	/*
		goos: linux
		goarch: amd64
		pkg: github.com/joshprzybyszewski/aoc2022
		cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz
		BenchmarkAll/Day_1/Part_One-4         	   18214	     68278 ns/op	   40983 B/op	       3 allocs/op
		BenchmarkAll/Day_1/Part_Two-4         	   14677	     83095 ns/op	   45096 B/op	      13 allocs/op
		BenchmarkAll/Day_2/Part_One-4         	    2995	    388317 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_2/Part_Two-4         	    3392	    404177 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_3/Part_One-4         	   21990	     54179 ns/op	    4868 B/op	       2 allocs/op
		BenchmarkAll/Day_3/Part_Two-4         	   20829	     56839 ns/op	   13780 B/op	      55 allocs/op
		BenchmarkAll/Day_4/Part_One-4         	    2812	    449865 ns/op	  145192 B/op	    3004 allocs/op
		BenchmarkAll/Day_4/Part_Two-4         	    2730	    449404 ns/op	  145190 B/op	    3004 allocs/op
		BenchmarkAll/Day_5/Part_One-4         	    1372	    894679 ns/op	   66394 B/op	    2615 allocs/op
		BenchmarkAll/Day_5/Part_Two-4         	    1353	    894593 ns/op	   66426 B/op	    2614 allocs/op
		BenchmarkAll/Day_6/Part_One-4         	  360046	      3374 ns/op	       4 B/op	       1 allocs/op
		BenchmarkAll/Day_6/Part_Two-4         	  404757	      2932 ns/op	       4 B/op	       1 allocs/op
		BenchmarkAll/Day_7/Part_One-4         	    4510	    235649 ns/op	   79863 B/op	    1811 allocs/op
		BenchmarkAll/Day_7/Part_Two-4         	    5328	    208960 ns/op	   76824 B/op	    1705 allocs/op
		PASS
		ok  	github.com/joshprzybyszewski/aoc2022	19.991s
	*/

	b.Run(`Day 1`, func(b *testing.B) {
		input, err := fetch.Input(1)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = one.One(input)
				if answer != `68787` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = one.Two(input)
				if answer != `198041` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 2`, func(b *testing.B) {
		input, err := fetch.Input(2)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = two.One(input)
				if answer != `11873` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = two.Two(input)
				if answer != `12014` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 3`, func(b *testing.B) {
		input, err := fetch.Input(3)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = three.One(input)
				if answer != `7863` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = three.Two(input)
				if answer != `2488` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 4`, func(b *testing.B) {
		input, err := fetch.Input(4)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = four.One(input)
				if answer != `500` {
					b.Logf("got answer: %q", answer)
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = four.Two(input)
				if answer != `815` {
					b.Logf("got answer: %q", answer)
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 5`, func(b *testing.B) {
		input, err := fetch.Input(5)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = five.One(input)
				if answer != `TWSGQHNHL` {
					b.Logf("Got answer: %q", answer)
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = five.Two(input)
				if answer != `JNRSCDWPP` {
					b.Logf("Got answer: %q", answer)
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 6`, func(b *testing.B) {
		input, err := fetch.Input(6)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = six.One(input)
				if answer != `1300` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = six.Two(input)
				if answer != `3986` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 7`, func(b *testing.B) {
		input, err := fetch.Input(7)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = seven.One(input)
				if answer != `1915606` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = seven.Two(input)
				if answer != `5025657` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})
}
