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
		BenchmarkAll/Day_1/Part_One-4         	   17542	     67235 ns/op	   40983 B/op	       3 allocs/op
		BenchmarkAll/Day_1/Part_Two-4         	   13522	     83610 ns/op	   45095 B/op	      13 allocs/op
		BenchmarkAll/Day_2/Part_One-4         	    3050	    399896 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_2/Part_Two-4         	    3580	    392836 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_3/Part_One-4         	   21208	     55440 ns/op	    4868 B/op	       2 allocs/op
		BenchmarkAll/Day_3/Part_Two-4         	   21242	     59418 ns/op	   13780 B/op	      55 allocs/op
		BenchmarkAll/Day_4/Part_One-4         	    3025	    450535 ns/op	  145190 B/op	    3004 allocs/op
		BenchmarkAll/Day_4/Part_Two-4         	    2763	    423436 ns/op	  145190 B/op	    3004 allocs/op
		BenchmarkAll/Day_5/Part_One-4         	    1252	    916973 ns/op	   66394 B/op	    2615 allocs/op
		BenchmarkAll/Day_5/Part_Two-4         	    1310	    889754 ns/op	   66427 B/op	    2614 allocs/op
		BenchmarkAll/Day_6/Part_One-4         	  353290	      3224 ns/op	       4 B/op	       1 allocs/op
		BenchmarkAll/Day_6/Part_Two-4         	  421429	      2957 ns/op	       4 B/op	       1 allocs/op
		BenchmarkAll/Day_7/Part_One-4         	    5552	    258292 ns/op	   91816 B/op	    1877 allocs/op
		BenchmarkAll/Day_7/Part_Two-4         	    5040	    249470 ns/op	   92335 B/op	    1871 allocs/op
		PASS
		ok  	github.com/joshprzybyszewski/aoc2022	21.510s
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
