package main

import (
	"testing"

	"github.com/joshprzybyszewski/aoc2022/eight"
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
		BenchmarkAll/Day_1/Part_One-4         	   18386	     66818 ns/op	   40984 B/op	       3 allocs/op
		BenchmarkAll/Day_1/Part_Two-4         	   14806	     82150 ns/op	   45096 B/op	      13 allocs/op
		BenchmarkAll/Day_2/Part_One-4         	    2860	    394875 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_2/Part_Two-4         	    2923	    381376 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_3/Part_One-4         	   21466	     55840 ns/op	    4868 B/op	       2 allocs/op
		BenchmarkAll/Day_3/Part_Two-4         	   20743	     60214 ns/op	   13780 B/op	      55 allocs/op
		BenchmarkAll/Day_4/Part_One-4         	    2518	    404115 ns/op	  112387 B/op	    3002 allocs/op
		BenchmarkAll/Day_4/Part_Two-4         	     793	   1475802 ns/op	   87342 B/op	    5460 allocs/op
		BenchmarkAll/Day_5/Part_One-4         	    1269	    896084 ns/op	   66394 B/op	    2615 allocs/op
		BenchmarkAll/Day_5/Part_Two-4         	    1369	    880554 ns/op	   66426 B/op	    2614 allocs/op
		BenchmarkAll/Day_6/Part_One-4         	  382370	      3056 ns/op	       4 B/op	       1 allocs/op
		BenchmarkAll/Day_6/Part_Two-4         	  407072	      2818 ns/op	       4 B/op	       1 allocs/op
		BenchmarkAll/Day_7/Part_One-4         	    5194	    229318 ns/op	   79862 B/op	    1811 allocs/op
		BenchmarkAll/Day_7/Part_Two-4         	    5654	    223778 ns/op	   76822 B/op	    1705 allocs/op
		BenchmarkAll/Day_8/Part_One-4         	    9171	    136176 ns/op	  106964 B/op	     202 allocs/op
		BenchmarkAll/Day_8/Part_Two-4         	    2378	    452643 ns/op	   93192 B/op	     102 allocs/op
		PASS
		ok  	github.com/joshprzybyszewski/aoc2022	23.078s
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
				if err != nil {
					b.Logf("got error: %q", err)
					b.Fail()
				}
				if answer != `500` {
					b.Logf("got answer: %q", answer)
					b.Fail()
				}

			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = four.Two(input)
				if err != nil {
					b.Logf("got error: %q", err)
					b.Fail()
				}
				if answer != `815` {
					b.Logf("got answer: %q", answer)
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

	b.Run(`Day 8`, func(b *testing.B) {
		input, err := fetch.Input(8)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = eight.One(input)
				if answer != `1809` {
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
				answer, err = eight.Two(input)
				if answer != `479400` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})
}
